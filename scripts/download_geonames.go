package main

import (
	"archive/zip"
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const geonamesURL = "https://download.geonames.org/export/dump/"

// adminToISO maps GeoNames administrative division codes to ISO 3166-2 subdivision codes.
// This is used during data generation to pre-compute ISO codes.
// The mapping is loaded from mapping.go at runtime.
var adminToISO map[string]string

func initAdminToISO() error {
	adminToISO = make(map[string]string)

	// Find mapping.go file (should be in parent directory)
	scriptDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Try to find mapping.go in the package root (parent of scripts/)
	mappingPath := filepath.Join(scriptDir, "..", "mapping.go")
	if _, err := os.Stat(mappingPath); os.IsNotExist(err) {
		// Try current directory
		mappingPath = filepath.Join(scriptDir, "mapping.go")
		if _, err := os.Stat(mappingPath); os.IsNotExist(err) {
			return fmt.Errorf("mapping.go not found")
		}
	}

	// Read and parse mapping.go
	data, err := os.ReadFile(mappingPath)
	if err != nil {
		return fmt.Errorf("failed to read mapping.go: %w", err)
	}

	// Parse map entries using regex: "KEY": "VALUE", // comment
	// Pattern: "AD.02": "AD-02", // Canillo
	re := regexp.MustCompile(`"([^"]+)":\s*"([^"]+)",`)
	matches := re.FindAllStringSubmatch(string(data), -1)
	for _, match := range matches {
		if len(match) == 3 {
			adminToISO[match[1]] = match[2]
		}
	}

	return nil
}

// Common countries to download
var countries = []string{
	"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	"AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI",
	"BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY",
	"BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN",
	"CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM",
	"DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK",
	"FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL",
	"GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM",
	"HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR",
	"IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN",
	"KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS",
	"LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK",
	"ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW",
	"MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	"NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM",
	"PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW",
	"SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM",
	"SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF",
	"TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW",
	"TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI",
	"VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW",
}

// GeonameEntry represents a geonames entry for lookup
// Collisions are resolved during generation (highest population wins)
// Population is not stored - only used during generation for collision resolution
type GeonameEntry struct {
	Name        string
	AsciiName   string
	Admin1Code  string
	CountryCode string
	ISOCode     string // Pre-computed ISO 3166-2 code (e.g., "DE-BY", "FR-ARA")
}

// getISOCode looks up the ISO 3166-2 code for a given admin key (e.g., "DE.02" -> "DE-BY")
// Returns empty string if not found
func getISOCode(adminKey string) string {
	if iso, ok := adminToISO[adminKey]; ok {
		return iso
	}
	return ""
}

// CityLookupMap is a map from lowercase city name to GeonameEntry
type CityLookupMap map[string]*GeonameEntry

// CountryLookupMap is a map from country code to city lookup map
type CountryLookupMap map[string]CityLookupMap

func main() {
	// Initialize adminToISO mapping
	if err := initAdminToISO(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to load adminToISO mapping: %v\n", err)
		fmt.Fprintf(os.Stderr, "ISO codes will not be pre-computed. Continuing anyway...\n")
	}

	// Get current working directory for debugging
	cwd, _ := os.Getwd()
	fmt.Printf("Current working directory: %s\n", cwd)

	// Create data/geonames directory relative to package root
	// Script can be run from package root or scripts/ directory
	zipDir := filepath.Join("data", "geonames")
	goDir := filepath.Join("data", "generated")

	absZipDir, _ := filepath.Abs(zipDir)
	absGoDir, _ := filepath.Abs(goDir)
	fmt.Printf("ZIP directory (relative): %s\n", zipDir)
	fmt.Printf("ZIP directory (absolute): %s\n", absZipDir)
	fmt.Printf("Go files directory (relative): %s\n", goDir)
	fmt.Printf("Go files directory (absolute): %s\n", absGoDir)

	if err := os.MkdirAll(zipDir, 0755); err != nil {
		fmt.Printf("Failed to create ZIP directory: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(goDir, 0755); err != nil {
		fmt.Printf("Failed to create Go directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloading and processing geonames data...\n")

	success := 0
	failed := 0

	for _, country := range countries {
		filename := country + ".zip"
		url := geonamesURL + filename
		localPath := filepath.Join(zipDir, filename)
		goPath := filepath.Join(goDir, country+".go")

		// Skip if Go file already exists
		if _, err := os.Stat(goPath); err == nil {
			fmt.Printf("✓ %s.go (already exists)\n", country)
			success++
			continue
		}

		// Download zip if it doesn't exist
		if _, err := os.Stat(localPath); err != nil {
			fmt.Printf("Downloading %s... ", filename)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("✗ download error: %v\n", err)
				failed++
				continue
			}

			if resp.StatusCode != 200 {
				fmt.Printf("✗ HTTP %d\n", resp.StatusCode)
				resp.Body.Close()
				failed++
				continue
			}

			out, err := os.Create(localPath)
			if err != nil {
				fmt.Printf("✗ create error: %v\n", err)
				resp.Body.Close()
				failed++
				continue
			}

			_, err = io.Copy(out, resp.Body)
			out.Close()
			resp.Body.Close()

			if err != nil {
				fmt.Printf("✗ copy error: %v\n", err)
				os.Remove(localPath)
				failed++
				continue
			}
		}

		// Process the zip file and create Go file
		fmt.Printf("Processing %s... ", filename)
		if err := processZipFile(localPath, goPath, country); err != nil {
			fmt.Printf("✗ processing error: %v\n", err)
			failed++
			continue
		}

		fmt.Printf("✓\n")
		success++
	}

	fmt.Printf("\nProcessing complete: %d successful, %d failed\n", success, failed)

	// Verify files were created and collect successfully generated countries
	fmt.Printf("\n=== Verification ===\n")
	files, err := filepath.Glob(filepath.Join(goDir, "*.go"))
	if err != nil {
		fmt.Printf("Error checking files: %v\n", err)
	} else {
		fmt.Printf("Found %d Go files in %s\n", len(files), goDir)
		if len(files) > 0 {
			fmt.Printf("First file: %s\n", files[0])
		}
	}

	// Collect successfully generated country files (excluding types.go and loader.go)
	var generatedCountries []string
	countryFiles, err := filepath.Glob(filepath.Join(goDir, "[A-Z][A-Z].go"))
	if err == nil {
		for _, file := range countryFiles {
			base := filepath.Base(file)
			countryCode := strings.TrimSuffix(base, ".go")
			if len(countryCode) == 2 {
				generatedCountries = append(generatedCountries, countryCode)
			}
		}
	}

	// Generate types file
	typesPath := filepath.Join(goDir, "types.go")
	fmt.Printf("\nGenerating types file...\n")
	if err := generateTypesFile(typesPath); err != nil {
		fmt.Printf("Failed to generate types: %v\n", err)
	} else {
		fmt.Printf("✓ Types file generated: %s\n", typesPath)
	}

	// Generate loader file with only successfully generated countries
	loaderPath := filepath.Join(goDir, "loader.go")
	fmt.Printf("Generating loader file with %d countries...\n", len(generatedCountries))
	if err := generateLoaderFile(loaderPath, generatedCountries); err != nil {
		fmt.Printf("Failed to generate loader: %v\n", err)
	} else {
		fmt.Printf("✓ Loader file generated: %s\n", loaderPath)
	}
}

// processZipFile extracts and processes a geonames zip file, creating a Go file with serialized data
func processZipFile(zipPath, goPath, countryCode string) error {
	// Open zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}
	defer r.Close()

	// Find the .txt file inside the zip (usually named like "AD.txt" or "US.txt")
	// Skip readme.txt and look for the country-specific file
	var txtFile *zip.File
	expectedTxtName := countryCode + ".txt"
	for _, f := range r.File {
		if f.Name == expectedTxtName {
			txtFile = f
			break
		}
	}

	// Fallback: if country-specific file not found, use any .txt file that's not readme.txt
	if txtFile == nil {
		for _, f := range r.File {
			if strings.HasSuffix(f.Name, ".txt") && !strings.HasPrefix(f.Name, "readme") {
				txtFile = f
				break
			}
		}
	}

	if txtFile == nil {
		return fmt.Errorf("no .txt file found in zip")
	}

	// Open the text file
	rc, err := txtFile.Open()
	if err != nil {
		return fmt.Errorf("failed to open text file: %w", err)
	}
	defer rc.Close()

	// Parse the geonames data
	// Format: geonameid, name, asciiname, alternatenames, latitude, longitude, feature class, feature code, country code, cc2, admin1 code, admin2 code, admin3 code, admin4 code, population, elevation, dem, timezone, modification date

	// Temporary structure to track population during processing
	type tempEntry struct {
		entry      *GeonameEntry
		population int
	}
	tempMap := make(map[string]*tempEntry)

	scanner := bufio.NewScanner(rc)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Split by tab
		fields := strings.Split(line, "\t")
		if len(fields) < 19 {
			continue // Skip malformed lines
		}

		// Extract relevant fields
		name := fields[1]
		asciiName := fields[2]
		alternateNames := fields[3] // Comma-separated list of alternative names
		admin1Code := fields[10]
		population := 0
		if fields[14] != "" {
			if p, err := strconv.Atoi(fields[14]); err == nil {
				population = p
			}
		}

		// Only process cities (feature class P = populated place)
		featureClass := fields[6]
		if featureClass != "P" {
			continue
		}

		// Compute ISO code from admin code
		adminKey := fmt.Sprintf("%s.%s", countryCode, admin1Code)
		isoCode := getISOCode(adminKey)

		// Create entry with pre-computed ISO code
		entry := &GeonameEntry{
			Name:        name,
			AsciiName:   asciiName,
			Admin1Code:  admin1Code,
			CountryCode: countryCode,
			ISOCode:     isoCode,
		}

		// Helper function to add entry to temp map, resolving collisions by population
		addToTempMap := func(key string) {
			keyLower := strings.ToLower(strings.TrimSpace(key))
			if keyLower == "" {
				return
			}
			if existing, exists := tempMap[keyLower]; !exists || population > existing.population {
				tempMap[keyLower] = &tempEntry{
					entry:      entry,
					population: population,
				}
			}
		}

		// Add primary name
		addToTempMap(name)

		// Add ascii name if different
		if asciiName != name {
			addToTempMap(asciiName)
		}

		// Parse and add alternative names
		if alternateNames != "" {
			altNames := strings.Split(alternateNames, ",")
			for _, altName := range altNames {
				altName = strings.TrimSpace(altName)
				if altName != "" && altName != name && altName != asciiName {
					addToTempMap(altName)
				}
			}
		}
	}

	// Convert temp map to final lookup map (without population)
	lookupMap := make(CityLookupMap)
	for key, temp := range tempMap {
		lookupMap[key] = temp.entry
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Create Go file with embedded data
	// Use the country code from the filename parameter, not from the data
	return createGoFile(goPath, countryCode, lookupMap)
}

// createGoFile creates a Go file with embedded serialized lookup map using go:embed
func createGoFile(goPath, countryCode string, lookupMap CityLookupMap) error {
	// Serialize the map to gob file
	gobPath := goPath + ".gob"
	file, err := os.Create(gobPath)
	if err != nil {
		return fmt.Errorf("failed to create gob file: %w", err)
	}

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(lookupMap); err != nil {
		file.Close()
		return fmt.Errorf("failed to encode data: %w", err)
	}
	file.Close()

	// Generate Go file that embeds the gob file
	goFile, err := os.Create(goPath)
	if err != nil {
		return fmt.Errorf("failed to create Go file: %w", err)
	}
	defer goFile.Close()

	// Write Go file content using go:embed
	fmt.Fprintf(goFile, "// Code generated by scripts/download_geonames.go. DO NOT EDIT.\n\n")
	fmt.Fprintf(goFile, "package generated\n\n")
	fmt.Fprintf(goFile, "import (\n")
	fmt.Fprintf(goFile, "\t\"bytes\"\n")
	fmt.Fprintf(goFile, "\t_ \"embed\"\n")
	fmt.Fprintf(goFile, "\t\"encoding/gob\"\n")
	fmt.Fprintf(goFile, "\t\"fmt\"\n")
	fmt.Fprintf(goFile, ")\n\n")

	// Embed the gob file using go:embed
	gobFileName := countryCode + ".go.gob"
	fmt.Fprintf(goFile, "//go:embed %s\n", gobFileName)
	fmt.Fprintf(goFile, "var %sDataBytes []byte\n\n", strings.ToLower(countryCode))

	// Function to load the data
	fmt.Fprintf(goFile, "func load%s() CityLookupMap {\n", strings.Title(strings.ToLower(countryCode)))
	fmt.Fprintf(goFile, "\tvar lookupMap CityLookupMap\n")
	fmt.Fprintf(goFile, "\tdecoder := gob.NewDecoder(bytes.NewReader(%sDataBytes))\n", strings.ToLower(countryCode))
	fmt.Fprintf(goFile, "\tif err := decoder.Decode(&lookupMap); err != nil {\n")
	fmt.Fprintf(goFile, "\t\tpanic(fmt.Sprintf(\"failed to decode %s geonames data: %%v\", err))\n", countryCode)
	fmt.Fprintf(goFile, "\t}\n")
	fmt.Fprintf(goFile, "\treturn lookupMap\n")
	fmt.Fprintf(goFile, "}\n")

	return nil
}

// generateTypesFile creates a types file with the lookup map types
func generateTypesFile(typesPath string) error {
	typesFile, err := os.Create(typesPath)
	if err != nil {
		return fmt.Errorf("failed to create types file: %w", err)
	}
	defer typesFile.Close()

	fmt.Fprintf(typesFile, "// Code generated by scripts/download_geonames.go. DO NOT EDIT.\n\n")
	fmt.Fprintf(typesFile, "package generated\n\n")
	fmt.Fprintf(typesFile, "// GeonameEntry represents a geonames entry for lookup\n")
	fmt.Fprintf(typesFile, "// Collisions are resolved during generation (highest population wins)\n")
	fmt.Fprintf(typesFile, "// Population is not stored - only used during generation for collision resolution\n")
	fmt.Fprintf(typesFile, "type GeonameEntry struct {\n")
	fmt.Fprintf(typesFile, "\tName        string\n")
	fmt.Fprintf(typesFile, "\tAsciiName   string\n")
	fmt.Fprintf(typesFile, "\tAdmin1Code  string\n")
	fmt.Fprintf(typesFile, "\tCountryCode string\n")
	fmt.Fprintf(typesFile, "\tISOCode     string // Pre-computed ISO 3166-2 code (e.g., \"DE-BY\", \"FR-ARA\")\n")
	fmt.Fprintf(typesFile, "}\n\n")
	fmt.Fprintf(typesFile, "// CityLookupMap is a map from lowercase city name to GeonameEntry\n")
	fmt.Fprintf(typesFile, "type CityLookupMap map[string]*GeonameEntry\n\n")
	fmt.Fprintf(typesFile, "// CountryLookupMap is a map from country code to city lookup map\n")
	fmt.Fprintf(typesFile, "type CountryLookupMap map[string]CityLookupMap\n")

	return nil
}

// generateLoaderFile creates a loader file that lazy-loads country data
func generateLoaderFile(loaderPath string, countries []string) error {
	loaderFile, err := os.Create(loaderPath)
	if err != nil {
		return fmt.Errorf("failed to create loader file: %w", err)
	}
	defer loaderFile.Close()

	fmt.Fprintf(loaderFile, "// Code generated by scripts/download_geonames.go. DO NOT EDIT.\n\n")
	fmt.Fprintf(loaderFile, "package generated\n\n")
	fmt.Fprintf(loaderFile, "import (\n")
	fmt.Fprintf(loaderFile, "\t\"sync\"\n")
	fmt.Fprintf(loaderFile, ")\n\n")

	// Create a map of country code to loader function
	fmt.Fprintf(loaderFile, "var countryLoaders = map[string]func() CityLookupMap{\n")
	for _, country := range countries {
		fmt.Fprintf(loaderFile, "\t%q: load%s,\n", country, strings.Title(strings.ToLower(country)))
	}
	fmt.Fprintf(loaderFile, "}\n\n")

	fmt.Fprintf(loaderFile, "var (\n")
	fmt.Fprintf(loaderFile, "\tlookupCache CountryLookupMap\n")
	fmt.Fprintf(loaderFile, "\tlookupMutex sync.RWMutex\n")
	fmt.Fprintf(loaderFile, ")\n\n")

	fmt.Fprintf(loaderFile, "// GetCountryLookup returns the city lookup map for a specific country, loading it on demand\n")
	fmt.Fprintf(loaderFile, "func GetCountryLookup(countryCode string) (CityLookupMap, bool) {\n")
	fmt.Fprintf(loaderFile, "\t// Check cache first (read lock)\n")
	fmt.Fprintf(loaderFile, "\tlookupMutex.RLock()\n")
	fmt.Fprintf(loaderFile, "\tif lookupCache == nil {\n")
	fmt.Fprintf(loaderFile, "\t\tlookupMutex.RUnlock()\n")
	fmt.Fprintf(loaderFile, "\t\tlookupMutex.Lock()\n")
	fmt.Fprintf(loaderFile, "\t\tif lookupCache == nil {\n")
	fmt.Fprintf(loaderFile, "\t\t\tlookupCache = make(CountryLookupMap)\n")
	fmt.Fprintf(loaderFile, "\t\t}\n")
	fmt.Fprintf(loaderFile, "\t\tlookupMutex.Unlock()\n")
	fmt.Fprintf(loaderFile, "\t\tlookupMutex.RLock()\n")
	fmt.Fprintf(loaderFile, "\t}\n")
	fmt.Fprintf(loaderFile, "\tif cached, exists := lookupCache[countryCode]; exists {\n")
	fmt.Fprintf(loaderFile, "\t\tlookupMutex.RUnlock()\n")
	fmt.Fprintf(loaderFile, "\t\treturn cached, true\n")
	fmt.Fprintf(loaderFile, "\t}\n")
	fmt.Fprintf(loaderFile, "\tlookupMutex.RUnlock()\n\n")

	fmt.Fprintf(loaderFile, "\t// Load on demand (write lock)\n")
	fmt.Fprintf(loaderFile, "\tloader, exists := countryLoaders[countryCode]\n")
	fmt.Fprintf(loaderFile, "\tif !exists {\n")
	fmt.Fprintf(loaderFile, "\t\treturn nil, false\n")
	fmt.Fprintf(loaderFile, "\t}\n\n")

	fmt.Fprintf(loaderFile, "\tlookupMutex.Lock()\n")
	fmt.Fprintf(loaderFile, "\tdefer lookupMutex.Unlock()\n\n")

	fmt.Fprintf(loaderFile, "\t// Double-check after acquiring write lock\n")
	fmt.Fprintf(loaderFile, "\tif cached, exists := lookupCache[countryCode]; exists {\n")
	fmt.Fprintf(loaderFile, "\t\treturn cached, true\n")
	fmt.Fprintf(loaderFile, "\t}\n\n")

	fmt.Fprintf(loaderFile, "\t// Load the country data\n")
	fmt.Fprintf(loaderFile, "\tlookupCache[countryCode] = loader()\n")
	fmt.Fprintf(loaderFile, "\treturn lookupCache[countryCode], true\n")
	fmt.Fprintf(loaderFile, "}\n")

	return nil
}
