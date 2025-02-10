package constants

// SearchResultTotalCountType indicates how the client should interpret the
// DataStoreSearchResult.totalCount value when counting search results
type SearchResultTotalCountType uint8

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
