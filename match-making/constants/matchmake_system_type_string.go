package constants

// MatchmakeSystemTypeString represents MatchmakeSystemType values but as strings.
// Used for MatchmakeSessionSearchCriteria, where the value is encoded as a string
// so that the field can be optional.
//
// Note: This is not a real NEX type, this is a supplementary type meant to make our
//
//	code easier to work with.
type MatchmakeSystemTypeString string

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
