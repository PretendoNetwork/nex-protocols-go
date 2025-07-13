package constants

// MatchmakeGeoIPResult represents an enum with an unknown use
type MatchmakeGeoIPResult uint32

// IsValid ensures the value of the MatchmakeGeoIPResult is within
// the expected range
func (mgipr MatchmakeGeoIPResult) IsValid() bool {
	return mgipr >= MatchmakeGeoIPResultInvalid && mgipr <= MatchmakeGeoIPResultUnused
}

const (
	// MatchmakeGeoIPResultInvalid indicates an invalid value
	MatchmakeGeoIPResultInvalid MatchmakeGeoIPResult = iota

	// MatchmakeGeoIPResultAnybody has an unknown use
	MatchmakeGeoIPResultFound

	// MatchmakeGeoIPResultFriends has an unknown use
	MatchmakeGeoIPResultNotFound

	// MatchmakeGeoIPResultFriendsInvite has an unknown use
	MatchmakeGeoIPResultUnused
)
