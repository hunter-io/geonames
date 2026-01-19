# geonames

A Go package for resolving city names to their administrative divisions (states/regions) using GeoNames data.

## Features

- Fast lazy-loading: Only loads country data when needed
- Alternative name support: Finds cities by their English names, local names, and alternatenames
- Pre-generated lookup maps: Serialized Go data files for fast access
- 249 countries supported

## Usage

```go
import "github.com/hunter-io/geonames"

result, err := geonames.ResolveState("DE", "Berlin")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("ISO Code: %s, State: %s\n", result.ISOCode, result.State)
// Output: ISO Code: DE-BE, State: 16
```

## Generating Data

To regenerate the data files from GeoNames:

```bash
go run scripts/download_geonames.go
```

This will:
1. Download GeoNames zip files for all countries
2. Extract and parse the data
3. Generate Go files with embedded binary data using `//go:embed`

## Structure

- `resolver.go` - Main API for resolving cities to states
- `city_alternatives.go` - City name variant handling
- `mapping.go` - GeoNames admin codes to ISO 3166-2 mapping
- `data/generated/` - Generated lookup maps (one per country)
- `scripts/download_geonames.go` - Data generation script

## Credits

This package uses data from [geonames.org](https://www.geonames.org/) and was inspired by the [mkrou/geonames](https://github.com/mkrou/geonames) Go parsing library.

## License

MIT