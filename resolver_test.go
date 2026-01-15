package geonames

import (
	"testing"
)

func TestResolveState(t *testing.T) {
	tests := []struct {
		name      string
		country   string
		city      string
		wantState string
		wantError bool
	}{
		// Test with countries that are currently generated (AD through BI)
		// Note: DE, FR, ES, IT, GB etc. will be added once those countries are generated
		{
			name:      "Vienna, Austria",
			country:   "AT",
			city:      "Vienna",
			wantState: "AT-9", // Vienna
			wantError: false,
		},
		{
			name:      "Non-existent city",
			country:   "AT",
			city:      "NonExistentCity",
			wantError: true,
		},
		{
			name:      "Empty country",
			country:   "",
			city:      "Vienna",
			wantError: true,
		},
		{
			name:      "Empty city",
			country:   "AT",
			city:      "",
			wantError: true,
		},
		{
			name:      "Country not in data",
			country:   "DE",
			city:      "Berlin",
			wantError: true, // DE not generated yet
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stateResult, err := ResolveState(tt.country, tt.city)

			if tt.wantError {
				if err == nil {
					t.Errorf("ResolveState() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("ResolveState() unexpected error: %v", err)
				return
			}

			if stateResult.ISOCode != tt.wantState {
				t.Errorf("ResolveState() = %v, want %v", stateResult.ISOCode, tt.wantState)
			}
		})
	}
}

func TestResolveState_CaseInsensitive(t *testing.T) {
	tests := []struct {
		country string
		city    string
	}{
		{"at", "vienna"},
		{"AT", "VIENNA"},
		{"At", "Vienna"},
	}

	for _, tt := range tests {
		t.Run(tt.country+"_"+tt.city, func(t *testing.T) {
			stateResult, err := ResolveState(tt.country, tt.city)
			if err != nil {
				// Skip if country not generated yet
				if err.Error() == "country '"+tt.country+"' not found in geonames data" {
					t.Skipf("Country %s not generated yet", tt.country)
				}
				t.Errorf("ResolveState() unexpected error: %v", err)
				return
			}

			// Vienna should be AT-9
			if stateResult.ISOCode != "AT-9" {
				t.Errorf("ResolveState() = %v, want AT-9", stateResult.ISOCode)
			}
		})
	}
}

// TestResolveState_Berlin will work once DE is generated
func TestResolveState_Berlin(t *testing.T) {
	stateResult, err := ResolveState("DE", "Berlin")
	if err != nil {
		if err.Error() == "country 'DE' not found in geonames data" {
			t.Skipf("DE not generated yet - run the full download script")
			return
		}
		t.Fatalf("ResolveState() unexpected error: %v", err)
	}

	if stateResult.ISOCode != "DE-BE" {
		t.Errorf("ResolveState() = %v, want DE-BE", stateResult.ISOCode)
	}
}

// TestResolveState_Munich will work once DE is generated
func TestResolveState_Munich(t *testing.T) {
	stateResult, err := ResolveState("DE", "Munich")
	if err != nil {
		if err.Error() == "country 'DE' not found in geonames data" {
			t.Skipf("DE not generated yet - run the full download script")
			return
		}
		t.Fatalf("ResolveState() unexpected error: %v", err)
	}

	if stateResult.ISOCode != "DE-BY" {
		t.Errorf("ResolveState() = %v, want DE-BY", stateResult.ISOCode)
	}
}
