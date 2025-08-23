package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// SearchSortOrder tells the server in what direction to order search results
type SearchSortOrder uint8

const (
	// SearchSortOrderAsc means the results should be ordered in ascending order
	SearchSortOrderAsc SearchSortOrder = iota

	// SearchSortOrderDesc means the results should be ordered in descending order
	SearchSortOrderDesc
)

// WriteTo writes the SearchSortOrder to the given writable
func (sso SearchSortOrder) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(sso))
}

// ExtractFrom extracts the SearchSortOrder value from the given readable
func (sso *SearchSortOrder) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*sso = SearchSortOrder(value)
	return nil
}
