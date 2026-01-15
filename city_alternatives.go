package geonames

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// cityNameAlternatives maps common English/international city names to their local variants
// This is a fallback for cases where geonames alternatenames are incomplete or missing.
// Most cities should be found via alternatenames indexing, but this provides additional coverage
// for cases where alternatenames are empty or don't include common English names.
// Key: lowercase English/international name
// Value: slice of local name variants (with proper diacritics/umlauts)
var cityNameAlternatives = map[string][]string{
	// German cities
	"cologne":   {"köln"},
	"munich":    {"münchen"},
	"nuremberg": {"nürnberg"},

	// Austrian cities
	"vienna": {"wien"},

	// Swiss cities
	"zurich": {"zürich"},
	"geneva": {"genève", "genf"},

	// Belgian cities
	"brussels": {"bruxelles", "brussel"},

	// Czech cities
	"prague": {"praha"},

	// Polish cities
	"warsaw": {"warszawa"},

	// Russian cities
	"moscow": {"moskva", "москва"},

	// Italian cities
	"rome":     {"roma"},
	"milan":    {"milano"},
	"florence": {"firenze"},
	"venice":   {"venezia"},
	"naples":   {"napoli"},

	// Portuguese cities
	"lisbon": {"lisboa"},

	// Danish cities
	"copenhagen": {"københavn"},

	// Greek cities
	"athens": {"athína", "αθήνα"},
}

// removeDiacritics removes accents and diacritics from a string
// Useful for normalizing city names (e.g., "Zürich" -> "Zurich")
func removeDiacritics(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// getCityNameVariants returns all possible name variants for a city
// This includes:
// - The original name
// - Known alternatives from the cityNameAlternatives map (fallback for missing alternatenames)
// - Version without diacritics (Zürich -> Zurich)
// - Version with common umlaut replacements (oe->ö, ae->ä, ue->ü)
// Note: Most variants should already be indexed via geonames alternatenames, but these
// transformations provide additional coverage for edge cases.
func getCityNameVariants(city string) []string {
	variants := []string{city}
	cityLower := strings.ToLower(city)

	// Add known alternatives from map (fallback for missing alternatenames)
	if alts, ok := cityNameAlternatives[cityLower]; ok {
		variants = append(variants, alts...)
	}

	// Add version without diacritics (Zürich -> Zurich)
	noDiacritics := removeDiacritics(city)
	if noDiacritics != city {
		variants = append(variants, noDiacritics)
	}

	// Add version with common umlaut replacements
	// This handles cases where users type "Koeln" instead of "Köln"
	if strings.Contains(cityLower, "oe") || strings.Contains(cityLower, "ae") || strings.Contains(cityLower, "ue") {
		withUmlauts := strings.ReplaceAll(cityLower, "oe", "ö")
		withUmlauts = strings.ReplaceAll(withUmlauts, "ae", "ä")
		withUmlauts = strings.ReplaceAll(withUmlauts, "ue", "ü")
		variants = append(variants, withUmlauts)
	}

	return variants
}
