package constants

// MatchmakeSessionOption0 has an unknown use.
// Seems related to whether or not the delay the response for "Auto"
// matchmaking methods, for some reason?
type MatchmakeSessionOption0 int64

const (
	// MatchmakeSessionOption0Random has an unknown use.
	MatchmakeSessionOption0None MatchmakeSessionOption0 = 0

	// MatchmakeSessionOption0ForceAutoMatchDelay has an unknown use.
	MatchmakeSessionOption0ForceAutomatchDelay MatchmakeSessionOption0 = 1

	// MatchmakeSessionOption0ForceAutomatchNoDelay has an unknown use.
	MatchmakeSessionOption0ForceAutomatchNoDelay MatchmakeSessionOption0 = 2
)
