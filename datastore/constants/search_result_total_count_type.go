package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// SearchResultTotalCountType indicates how the client should interpret the
// DataStoreSearchResult.totalCount value when counting search results
type SearchResultTotalCountType uint8

// WriteTo writes the SearchResultTotalCountType to the given writable
func (srtct SearchResultTotalCountType) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(srtct))
}

// ExtractFrom extracts the SearchResultTotalCountType value from the given readable
func (srtct *SearchResultTotalCountType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*srtct = SearchResultTotalCountType(value)
	return nil
}

const (
	// SearchResultTotalExact means the returned count is the
	// exact number of objects which match the search criteria
	SearchResultTotalExact SearchResultTotalCountType = iota

	// SearchResultTotalMinimum means the returned count is the
	// mininum number of objects which match the search criteria.
	// There are more objects beyond this count
	SearchResultTotalMinimum

	// SearchResultTotalEstimate means the returned count is an
	// estimate of the number of objects which match the search
	// criteria
	SearchResultTotalEstimate

	// SearchResultTotalDisabled means there is no returned count
	SearchResultTotalDisabled
)
