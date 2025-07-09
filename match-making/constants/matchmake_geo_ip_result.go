package constants

// MatchmakeGeoIPResult represents an enum with an unknown use
type MatchmakeGeoIPResult uint32

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
