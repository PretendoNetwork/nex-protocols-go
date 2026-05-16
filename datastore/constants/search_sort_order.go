package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SearchSortOrder tells the server in what direction to order search results
type SearchSortOrder uint8

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

// String returns a human-readable representation of the SearchSortOrder.
func (sso SearchSortOrder) String() string {
	switch sso {
	case SearchSortOrderAsc:
		return "Asc"
	case SearchSortOrderDesc:
		return "Desc"
	default:
		return fmt.Sprintf("SearchSortOrder(%d)", int(sso))
	}
}

const (
	// SearchSortOrderAsc means the results should be ordered in ascending order
	SearchSortOrderAsc SearchSortOrder = iota

	// SearchSortOrderDesc means the results should be ordered in descending order
	SearchSortOrderDesc
)
