package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// SearchTarget represents the type of user who owns an object.
// Not to be confused with DataStoreSearchParam.searchTarget,
// this is actually stored in DataStoreSearchParam.ownertype.
// Used to narrow search results based on owner type
type SearchTarget uint8

// WriteTo writes the SearchTarget to the given writable
func (st SearchTarget) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(st))
}

// ExtractFrom extracts the SearchTarget value from the given readable
func (st *SearchTarget) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*st = SearchTarget(value)
	return nil
}

const (
	// SearchTargetAnybody selects objects owned by anyone
	SearchTargetAnybody SearchTarget = iota

	// SearchTargetFriend selects objects owned by the users friends
	SearchTargetFriend

	// SearchTargetAnybodyExcludeSpecified selects objects owned by anyone
	// EXCEPT those set in DataStoreSearchParam.ownerIds
	SearchTargetAnybodyExcludeSpecified
)
