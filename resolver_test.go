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
		{
			name:      "Berlin, Germany",
			country:   "DE",
			city:      "Berlin",
			wantState: "DE-BE",
			wantError: false,
		},
		// Additional German cities
		{
			name:      "Hamburg, Germany",
			country:   "DE",
			city:      "Hamburg",
			wantState: "DE-HH",
			wantError: false,
		},
		{
			name:      "Munich, Germany (Bavaria)",
			country:   "DE",
			city:      "Munich",
			wantState: "DE-BY",
			wantError: false,
		},
		{
			name:      "Frankfurt am Main, Germany (Hesse)",
			country:   "DE",
			city:      "Frankfurt am Main",
			wantState: "DE-HE",
			wantError: false,
		},
		{
			name:      "Stuttgart, Germany (Baden-Württemberg)",
			country:   "DE",
			city:      "Stuttgart",
			wantState: "DE-BW",
			wantError: false,
		},
		{
			name:      "Düsseldorf, Germany (North Rhine-Westphalia)",
			country:   "DE",
			city:      "Düsseldorf",
			wantState: "DE-NW",
			wantError: false,
		},
		{
			name:      "Dortmund, Germany (North Rhine-Westphalia)",
			country:   "DE",
			city:      "Dortmund",
			wantState: "DE-NW",
			wantError: false,
		},
		{
			name:      "Essen, Germany (North Rhine-Westphalia)",
			country:   "DE",
			city:      "Essen",
			wantState: "DE-NW",
			wantError: false,
		},
		{
			name:      "Leipzig, Germany (Saxony)",
			country:   "DE",
			city:      "Leipzig",
			wantState: "DE-SN",
			wantError: false,
		},
		{
			name:      "Dresden, Germany (Saxony)",
			country:   "DE",
			city:      "Dresden",
			wantState: "DE-SN",
			wantError: false,
		},
		{
			name:      "Hannover, Germany (Lower Saxony)",
			country:   "DE",
			city:      "Hannover",
			wantState: "DE-NI",
			wantError: false,
		},
		{
			name:      "Paris, France",
			country:   "FR",
			city:      "Paris",
			wantState: "FR-IDF",
			wantError: false,
		},
		// Additional French cities
		{
			name:      "Marseille, France (Provence-Alpes-Côte d'Azur)",
			country:   "FR",
			city:      "Marseille",
			wantState: "FR-PAC",
			wantError: false,
		},
		{
			name:      "Lyon, France (Auvergne-Rhône-Alpes)",
			country:   "FR",
			city:      "Lyon",
			wantState: "FR-ARA",
			wantError: false,
		},
		{
			name:      "Toulouse, France (Occitanie)",
			country:   "FR",
			city:      "Toulouse",
			wantState: "FR-OCC",
			wantError: false,
		},
		{
			name:      "Nice, France (Provence-Alpes-Côte d'Azur)",
			country:   "FR",
			city:      "Nice",
			wantState: "FR-PAC",
			wantError: false,
		},
		{
			name:      "Nantes, France (Pays de la Loire)",
			country:   "FR",
			city:      "Nantes",
			wantState: "FR-PDL",
			wantError: false,
		},
		{
			name:      "Strasbourg, France (Grand Est)",
			country:   "FR",
			city:      "Strasbourg",
			wantState: "FR-GES",
			wantError: false,
		},
		{
			name:      "Montpellier, France (Occitanie)",
			country:   "FR",
			city:      "Montpellier",
			wantState: "FR-OCC",
			wantError: false,
		},
		{
			name:      "Bordeaux, France (Nouvelle-Aquitaine)",
			country:   "FR",
			city:      "Bordeaux",
			wantState: "FR-NAQ",
			wantError: false,
		},
		{
			name:      "Lille, France (Hauts-de-France)",
			country:   "FR",
			city:      "Lille",
			wantState: "FR-HDF",
			wantError: false,
		},
		{
			name:      "Rennes, France (Brittany)",
			country:   "FR",
			city:      "Rennes",
			wantState: "FR-BRE",
			wantError: false,
		},
		{
			name:      "Barcelona, Spain",
			country:   "ES",
			city:      "Barcelona",
			wantState: "ES-CT",
			wantError: false,
		},
		// Additional Spanish cities
		{
			name:      "Seville, Spain (Andalusia)",
			country:   "ES",
			city:      "Seville",
			wantState: "ES-AN",
			wantError: false,
		},
		{
			name:      "Zaragoza, Spain (Aragon)",
			country:   "ES",
			city:      "Zaragoza",
			wantState: "ES-AR",
			wantError: false,
		},
		{
			name:      "Málaga, Spain (Andalusia)",
			country:   "ES",
			city:      "Málaga",
			wantState: "ES-AN",
			wantError: false,
		},
		{
			name:      "Murcia, Spain (Region of Murcia)",
			country:   "ES",
			city:      "Murcia",
			wantState: "ES-MU",
			wantError: false,
		},
		{
			name:      "Palma, Spain (Balearic Islands)",
			country:   "ES",
			city:      "Palma",
			wantState: "ES-IB",
			wantError: false,
		},
		{
			name:      "Las Palmas, Spain (Canary Islands)",
			country:   "ES",
			city:      "Las Palmas",
			wantState: "ES-CN",
			wantError: false,
		},
		{
			name:      "Bilbao, Spain (Basque Country)",
			country:   "ES",
			city:      "Bilbao",
			wantState: "ES-PV",
			wantError: false,
		},
		{
			name:      "Alicante, Spain (Valencian Community)",
			country:   "ES",
			city:      "Alicante",
			wantState: "ES-VC",
			wantError: false,
		},
		{
			name:      "Rome, Italy",
			country:   "IT",
			city:      "Rome",
			wantState: "IT-62",
			wantError: false,
		},
		{
			name:      "Vilnius, Lithuania",
			country:   "LT",
			city:      "Vilnius",
			wantState: "LT-VL",
			wantError: false,
		},
		{
			name:      "London, United Kingdom",
			country:   "GB",
			city:      "London",
			wantState: "GB-ENG", // England
			wantError: false,
		},
		// Additional UK cities
		{
			name:      "Manchester, United Kingdom (England)",
			country:   "GB",
			city:      "Manchester",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Birmingham, United Kingdom (England)",
			country:   "GB",
			city:      "Birmingham",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Glasgow, United Kingdom (Scotland)",
			country:   "GB",
			city:      "Glasgow",
			wantState: "GB-SCT",
			wantError: false,
		},
		{
			name:      "Edinburgh, United Kingdom (Scotland)",
			country:   "GB",
			city:      "Edinburgh",
			wantState: "GB-SCT",
			wantError: false,
		},
		{
			name:      "Liverpool, United Kingdom (England)",
			country:   "GB",
			city:      "Liverpool",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Bristol, United Kingdom (England)",
			country:   "GB",
			city:      "Bristol",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Leeds, United Kingdom (England)",
			country:   "GB",
			city:      "Leeds",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Sheffield, United Kingdom (England)",
			country:   "GB",
			city:      "Sheffield",
			wantState: "GB-ENG",
			wantError: false,
		},
		{
			name:      "Cardiff, United Kingdom (Wales)",
			country:   "GB",
			city:      "Cardiff",
			wantState: "GB-WLS",
			wantError: false,
		},
		{
			name:      "Madrid, Spain",
			country:   "ES",
			city:      "Madrid",
			wantState: "ES-MD", // Community of Madrid
			wantError: false,
		},
		{
			name:      "Lisbon, Portugal",
			country:   "PT",
			city:      "Lisbon",
			wantState: "PT-11", // Lisboa
			wantError: false,
		},
		{
			name:      "Stockholm, Sweden",
			country:   "SE",
			city:      "Stockholm",
			wantState: "SE-AB", // Stockholm County
			wantError: false,
		},
		{
			name:      "Helsinki, Finland",
			country:   "FI",
			city:      "Helsinki",
			wantState: "FI-18", // Uusimaa
			wantError: false,
		},
		{
			name:      "Dublin, Ireland",
			country:   "IE",
			city:      "Dublin",
			wantState: "IE-L", // Leinster
			wantError: false,
		},
		{
			name:      "Bucharest, Romania",
			country:   "RO",
			city:      "Bucharest",
			wantState: "RO-B", // București
			wantError: false,
		},
		{
			name:      "Budapest, Hungary",
			country:   "HU",
			city:      "Budapest",
			wantState: "HU-BU", // Budapest
			wantError: false,
		},
		{
			name:      "Sofia, Bulgaria",
			country:   "BG",
			city:      "Sofia",
			wantState: "BG-22", // Sofia City
			wantError: false,
		},
		{
			name:      "Zagreb, Croatia",
			country:   "HR",
			city:      "Zagreb",
			wantState: "HR-21", // Zagreb City
			wantError: false,
		},
		{
			name:      "Bratislava, Slovakia",
			country:   "SK",
			city:      "Bratislava",
			wantState: "SK-BL", // Bratislavský kraj
			wantError: false,
		},
		{
			name:      "Ljubljana, Slovenia",
			country:   "SI",
			city:      "Ljubljana",
			wantState: "SI-061", // Ljubljana
			wantError: false,
		},
		{
			name:      "Tallinn, Estonia",
			country:   "EE",
			city:      "Tallinn",
			wantState: "EE-37", // Harju maakond
			wantError: false,
		},
		{
			name:      "Oslo, Norway",
			country:   "NO",
			city:      "Oslo",
			wantState: "NO-03", // Oslo
			wantError: false,
		},
		{
			name:      "Bern, Switzerland",
			country:   "CH",
			city:      "Bern",
			wantState: "CH-BE", // Bern
			wantError: false,
		},
		// Additional Swiss cities
		{
			name:      "Zürich, Switzerland",
			country:   "CH",
			city:      "Zürich",
			wantState: "CH-ZH",
			wantError: false,
		},
		{
			name:      "Geneva, Switzerland",
			country:   "CH",
			city:      "Geneva",
			wantState: "CH-GE",
			wantError: false,
		},
		{
			name:      "Basel, Switzerland",
			country:   "CH",
			city:      "Basel",
			wantState: "CH-BS",
			wantError: false,
		},
		{
			name:      "Lausanne, Switzerland",
			country:   "CH",
			city:      "Lausanne",
			wantState: "CH-VD",
			wantError: false,
		},
		{
			name:      "Krakow, Poland (Małopolskie)",
			country:   "PL",
			city:      "Krakow",
			wantState: "PL-12", // Lesser Poland
			wantError: false,
		},
		{
			name:      "Valencia, Spain (Valencian Community)",
			country:   "ES",
			city:      "Valencia",
			wantState: "ES-VC", // Valencian Community
			wantError: false,
		},
		{
			name:      "Cologne, Germany (North Rhine-Westphalia)",
			country:   "DE",
			city:      "Cologne",
			wantState: "DE-NW", // Nordrhein-Westfalen - English vs German name
			wantError: false,
		},
		{
			name:      "Non-existent city",
			country:   "DE",
			city:      "NonExistentCity",
			wantError: true,
		},
		{
			name:      "Empty country",
			country:   "",
			city:      "Berlin",
			wantError: true,
		},
		{
			name:      "Empty city",
			country:   "DE",
			city:      "",
			wantError: true,
		},
		{
			name:      "Country not in data",
			country:   "XX", // Non-existent country code
			city:      "City",
			wantError: true,
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
		{"de", "berlin"},
		{"DE", "BERLIN"},
		{"De", "Berlin"},
	}

	for _, tt := range tests {
		t.Run(tt.country+"_"+tt.city, func(t *testing.T) {
			stateResult, err := ResolveState(tt.country, tt.city)
			if err != nil {
				t.Errorf("ResolveState() unexpected error: %v", err)
				return
			}

			if stateResult.ISOCode != "DE-BE" {
				t.Errorf("ResolveState() = %v, want DE-BE", stateResult.ISOCode)
			}
		})
	}
}
