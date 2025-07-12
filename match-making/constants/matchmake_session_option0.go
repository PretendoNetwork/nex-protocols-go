package constants

// MatchmakeSessionOption0 has an unknown use.
// Seems related to whether or not the delay the response for "Auto"
// matchmaking methods, for some reason?
type MatchmakeSessionOption0 int64

func (mso MatchmakeSessionOption0) HasFlag(flag MatchmakeSessionOption0) bool {
	return mso&flag == flag
}

func (mso MatchmakeSessionOption0) HasFlags(flags ...MatchmakeSessionOption0) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if mso&flag != flag {
			return false
		}
	}

	return true
}

const (
	// MatchmakeSessionOption0None has an unknown use.
	MatchmakeSessionOption0None MatchmakeSessionOption0 = 0

	// MatchmakeSessionOption0ForceAutoMatchDelay has an unknown use.
	MatchmakeSessionOption0ForceAutomatchDelay MatchmakeSessionOption0 = 1

	// MatchmakeSessionOption0ForceAutomatchNoDelay has an unknown use.
	MatchmakeSessionOption0ForceAutomatchNoDelay MatchmakeSessionOption0 = 2
)
