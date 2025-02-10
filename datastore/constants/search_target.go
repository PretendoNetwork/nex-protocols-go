package constants

// SearchTarget represents the type of user who owns an object.
// Not to be confused with DataStoreSearchParam.searchTarget,
// this is actually stored in DataStoreSearchParam.ownertype.
// Used to narrow search results based on owner type
type SearchTarget uint8

const (
	// SearchTargetAnybody selects objects owned by anyone
	SearchTargetAnybody SearchTarget = iota

	// SearchTargetFriend selects objects owned by the users friends
	SearchTargetFriend

	// SearchTargetAnybodyExcludeSpecified selects objects owned by anyone
	// EXCEPT those set in DataStoreSearchParam.ownerIds
	SearchTargetAnybodyExcludeSpecified
)
