package constants

// SearchSortOrder tells the server in what direction to order search results
type SearchSortOrder uint8

const (
	// SearchSortOrderAsc means the results should be ordered in ascending order
	SearchSortOrderAsc SearchSortOrder = iota

	// SearchSortOrderDesc means the results should be ordered in descending order
	SearchSortOrderDesc
)
