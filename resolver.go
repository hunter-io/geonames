package geonames

import (
	"fmt"
	"strings"

	"github.com/hunter-io/geonames/data/generated"
)

//go:generate go run scripts/download_geonames.go

// StateResult contains both the ISO 3166-2 code and the state code
type StateResult struct {
	ISOCode string // Full ISO 3166-2 format (e.g., "DE-BY", "FR-ARA")
	State   string // Just the state/region code (e.g., "BY", "ARA")
}

// ResolveState resolves a city to its state/region
// country: ISO 3166-1 alpha-2 country code (e.g., "DE", "FR", "LT")
// city: City name to look up
// Returns StateResult with both ISO 3166-2 format and state code, or an error if not found
func ResolveState(country, city string) (StateResult, error) {
	// Normalize inputs
	country = strings.ToUpper(strings.TrimSpace(country))
	city = strings.TrimSpace(city)

	if country == "" || city == "" {
		return StateResult{}, fmt.Errorf("country and city are required")
	}

	// Get the country-specific lookup map (lazy-loaded)
	countryLookup, ok := generated.GetCountryLookup(country)
	if !ok {
		return StateResult{}, fmt.Errorf("country '%s' not found in geonames data", country)
	}

	var entry *generated.GeonameEntry
	var found bool

	// Get possible name variants (English name, local name with umlauts, etc.)
	cityVariants := getCityNameVariants(city)

	// Check each variant against the lookup map
	// Collisions are already resolved during generation (highest population wins)
	for _, variant := range cityVariants {
		variantLower := strings.ToLower(variant)
		if e, exists := countryLookup[variantLower]; exists {
			entry = e
			found = true
			break
		}
	}

	if !found || entry == nil {
		return StateResult{}, fmt.Errorf("city '%s' not found in country '%s'", city, country)
	}

	if entry.ISOCode == "" {
		return StateResult{}, fmt.Errorf("ISO code not found for city '%s' in country '%s'", city, country)
	}

	// Extract state code from ISOCode (e.g., "DE-BY" -> "BY")
	state := ""
	if entry.ISOCode != "" {
		parts := strings.Split(entry.ISOCode, "-")
		if len(parts) == 2 {
			state = parts[1]
		}
	}

	return StateResult{
		ISOCode: entry.ISOCode,
		State:   state,
	}, nil
}
