package constants

// MatchmakeSessionOption0 has an unknown use.
// Seems related to whether or not the delay the response for "Auto"
// matchmaking methods, for some reason?
type MatchmakeSessionOption0 uint32

const (
	// MatchmakeSessionOption0Random has an unknown use.
	MatchmakeSessionOption0None MatchmakeSessionOption0 = iota

	// MatchmakeSessionOption0ForceAutoMatchDelay has an unknown use.
	MatchmakeSessionOption0ForceAutoMatchDelay

	// MatchmakeSessionOption0AutoMatchNoDelay has an unknown use.
	MatchmakeSessionOption0ForceAutoMatchNoDelay
)
