package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// MatchmakeSystemTypeString represents MatchmakeSystemType values but as strings.
// Used for MatchmakeSessionSearchCriteria, where the value is encoded as a string
// so that the field can be optional.
//
// Note: This is not a real NEX type, this is a supplementary type meant to make our
//
//	code easier to work with.
type MatchmakeSystemTypeString string

// WriteTo writes the MatchmakeSystemTypeString to the given writable
func (msts MatchmakeSystemTypeString) WriteTo(writable types.Writable) {
	types.String(msts).WriteTo(writable)
}

// ExtractFrom extracts the MatchmakeSystemTypeString value from the given readable
func (msts *MatchmakeSystemTypeString) ExtractFrom(readable types.Readable) error {
	var s types.String
	if err := s.ExtractFrom(readable); err != nil {
		return err
	}
	*msts = MatchmakeSystemTypeString(s)
	return nil
}

const (
	// MatchmakeSystemTypeStringInvalid indicates that the field in MatchmakeSessionSearchCriteria
	// should be skipped.
	MatchmakeSystemTypeStringMissing MatchmakeSystemTypeString = ""

	// MatchmakeSystemTypeStringInvalid represents MatchmakeSystemTypeInvalid but as a string.
	MatchmakeSystemTypeStringInvalid MatchmakeSystemTypeString = "0"

	// MatchmakeSystemTypeStringAnybody represents MatchmakeSystemTypeAnybody but as a string.
	MatchmakeSystemTypeStringAnybody MatchmakeSystemTypeString = "1"

	// MatchmakeSystemTypeStringFriends represents MatchmakeSystemTypeFriends but as a string.
	MatchmakeSystemTypeStringFriends MatchmakeSystemTypeString = "2"

	// MatchmakeSystemTypeStringFriendsInvite represents MatchmakeSystemTypeFriendsInvite but as a string.
	MatchmakeSystemTypeStringFriendsInvite MatchmakeSystemTypeString = "3"

	// MatchmakeSystemTypeStringFriends represents MatchmakeSystemTypeInvite but as a string.
	MatchmakeSystemTypeStringInvite MatchmakeSystemTypeString = "4"

	// MatchmakeSystemTypeStringFriends represents MatchmakeSystemTypePersistentGathering but as a string.
	MatchmakeSystemTypeStringPersistentGathering MatchmakeSystemTypeString = "5"
)
