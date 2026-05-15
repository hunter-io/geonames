package geonames

import "testing"

func TestIsValidSubdivision(t *testing.T) {
	tests := []struct {
		name    string
		country string
		state   string
		want    bool
	}{
		// Valid ISO 3166-2 subdivisions (fully qualified).
		{name: "US California (prefixed)", country: "US", state: "US-CA", want: true},
		{name: "DE Bavaria (prefixed)", country: "DE", state: "DE-BY", want: true},
		{name: "FR Île-de-France (prefixed)", country: "FR", state: "FR-IDF", want: true},
		{name: "GB England (prefixed)", country: "GB", state: "GB-ENG", want: true},

		// Valid bare-state inputs (function auto-prefixes).
		{name: "US California (bare)", country: "US", state: "CA", want: true},
		{name: "DE Bavaria (bare)", country: "DE", state: "BY", want: true},
		{name: "FR Île-de-France (bare)", country: "FR", state: "IDF", want: true},

		// Case-insensitive.
		{name: "lowercase country and state", country: "fr", state: "idf", want: true},
		{name: "mixed case", country: "Us", state: "ca", want: true},

		// Whitespace tolerated.
		{name: "trailing whitespace", country: "FR ", state: " FR-IDF ", want: true},

		// FR INSEE numeric codes — not valid ISO, not produced by geonames.
		{name: "FR INSEE 75 (Paris dept)", country: "FR", state: "75", want: false},
		{name: "FR INSEE 20 (Corsica dept)", country: "FR", state: "20", want: false},
		{name: "FR INSEE 977 (Saint-Barthélemy dept)", country: "FR", state: "977", want: false},

		// US state code paired with wrong country.
		{name: "US code on FR", country: "FR", state: "LA", want: false},
		{name: "US code on FR (prefixed)", country: "FR", state: "FR-LA", want: false},
		{name: "US code on DE", country: "DE", state: "NY", want: false},

		// Unknown countries.
		{name: "unknown country code", country: "XX", state: "ANY", want: false},
		{name: "empty country", country: "", state: "CA", want: false},
		{name: "empty state", country: "US", state: "", want: false},
		{name: "both empty", country: "", state: "", want: false},

		// Junk subdivision codes.
		{name: "garbage state for valid country", country: "US", state: "ZZ", want: false},
		{name: "garbage state with prefix", country: "FR", state: "FR-NOTREAL", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSubdivision(tt.country, tt.state); got != tt.want {
				t.Errorf("IsValidSubdivision(%q, %q) = %v, want %v",
					tt.country, tt.state, got, tt.want)
			}
		})
	}
}

func TestIsValidSubdivision_CacheConsistency(t *testing.T) {
	// First call populates the cache; second call hits it. Both should return
	// the same answer.
	if !IsValidSubdivision("FR", "FR-IDF") {
		t.Fatal("first call: expected FR-IDF to be valid")
	}
	if !IsValidSubdivision("FR", "FR-IDF") {
		t.Fatal("second call (cached): expected FR-IDF to be valid")
	}
	if IsValidSubdivision("FR", "FR-LA") {
		t.Fatal("expected FR-LA to be invalid on cached country")
	}
}

func TestSubdivisionsOf(t *testing.T) {
	set, ok := SubdivisionsOf("FR")
	if !ok {
		t.Fatal("expected FR to be present in geonames data")
	}
	if !set["FR-IDF"] {
		t.Error("expected FR-IDF in FR subdivision set")
	}
	if set["FR-LA"] {
		t.Error("did not expect FR-LA (US code) in FR subdivision set")
	}
	if set["75"] {
		t.Error("did not expect raw INSEE 75 in FR subdivision set")
	}

	if _, ok := SubdivisionsOf("XX"); ok {
		t.Error("expected XX to be absent from geonames data")
	}
	if _, ok := SubdivisionsOf(""); ok {
		t.Error("expected empty country to return false")
	}
}
