package constants

// MatchmakeSystemType represents the method of matchmaking being used
type MatchmakeSystemType uint32

// IsValid ensures the value of the MatchmakeSystemType is within
// the expected range
func (mst MatchmakeSystemType) IsValid() bool {
	return mst >= MatchmakeSystemTypeInvalid && mst <= MatchmakeSystemTypePersistentGathering
}

const (
	// MatchmakeSystemTypeInvalid indicates an invalid value
	MatchmakeSystemTypeInvalid MatchmakeSystemType = iota

	// MatchmakeSystemTypeAnybody indicates that anybody can join the MatchmakeSession
	MatchmakeSystemTypeAnybody

	// MatchmakeSystemTypeFriends indicates that only friends of the owner can join the MatchmakeSession
	MatchmakeSystemTypeFriends

	// MatchmakeSystemTypeFriendsInvite indicates that only friends of the owner with invitation can join the MatchmakeSession
	MatchmakeSystemTypeFriendsInvite

	// MatchmakeSystemTypeFriends indicates that anybody with an invitation can join the MatchmakeSession
	MatchmakeSystemTypeInvite

	// MatchmakeSystemTypeFriends indicates that the MatchmakeSession is part of a PersistentGathering
	MatchmakeSystemTypePersistentGathering
)
