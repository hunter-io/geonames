package geonames

import (
	"strings"
	"sync"

	"github.com/hunter-io/geonames/data/generated"
)

// IsValidSubdivision reports whether (country, state) is a recognized ISO 3166-2
// subdivision in the geonames dataset.
//
// "Recognized" means: at least one city in geonames maps to this state code.
// A code that no city resolves to (e.g. an INSEE-only French numeric like "75",
// or a US state code paired with a non-US country) returns false.
//
// state may be supplied either bare ("CA") or fully qualified ("US-CA"); both
// resolve to the canonical "<country>-<state>" form before comparison. Inputs
// are case-insensitive.
//
// The validation set is derived once per country from the underlying city
// lookup map and cached.
func IsValidSubdivision(country, state string) bool {
	country = strings.ToUpper(strings.TrimSpace(country))
	state = strings.ToUpper(strings.TrimSpace(state))
	if country == "" || state == "" {
		return false
	}

	canonical := state
	prefix := country + "-"
	if !strings.HasPrefix(canonical, prefix) {
		canonical = prefix + canonical
	}

	set, ok := subdivisionSet(country)
	if !ok {
		return false
	}
	return set[canonical]
}

// SubdivisionsOf returns the set of valid ISO 3166-2 codes for the given
// country, derived from the geonames city data. Returns nil and false if the
// country is not present in the dataset.
//
// Useful for callers that need to enumerate (e.g. building UI dropdowns) or
// validate many states in a hot path without paying the lookup-and-prefix cost
// per call.
func SubdivisionsOf(country string) (map[string]bool, bool) {
	country = strings.ToUpper(strings.TrimSpace(country))
	if country == "" {
		return nil, false
	}
	return subdivisionSet(country)
}

var (
	subdivisionCache   = make(map[string]map[string]bool)
	subdivisionCacheMu sync.RWMutex
)

// subdivisionSet returns the set of canonical ISO 3166-2 codes that appear in
// the geonames data for the given country. Computed lazily, then cached.
func subdivisionSet(country string) (map[string]bool, bool) {
	subdivisionCacheMu.RLock()
	if set, ok := subdivisionCache[country]; ok {
		subdivisionCacheMu.RUnlock()
		return set, true
	}
	subdivisionCacheMu.RUnlock()

	lookup, ok := generated.GetCountryLookup(country)
	if !ok {
		return nil, false
	}

	set := make(map[string]bool)
	for _, entry := range lookup {
		if entry == nil || entry.ISOCode == "" {
			continue
		}
		set[entry.ISOCode] = true
	}

	subdivisionCacheMu.Lock()
	subdivisionCache[country] = set
	subdivisionCacheMu.Unlock()

	return set, true
}
