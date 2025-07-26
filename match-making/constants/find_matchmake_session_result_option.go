package constants

// FindMatchmakeSessionResultOption indicates how to populate the
// responses from FindMatchmakeSessionByParticipant
type FindMatchmakeSessionResultOption uint32

func (fmsr FindMatchmakeSessionResultOption) HasFlag(flag FindMatchmakeSessionResultOption) bool {
	return fmsr&flag == flag
}

func (fmsr FindMatchmakeSessionResultOption) HasFlags(flags ...FindMatchmakeSessionResultOption) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if fmsr&flag != flag {
			return false
		}
	}

	return true
}

const (
	// FindMatchmakeSessionResultOptionNone indicates no options
	FindMatchmakeSessionResultOptionNone FindMatchmakeSessionResultOption = 0

	// FindMatchmakeSessionResultOptionApplicationBuffer populates m_ApplicationBuffer in the results
	FindMatchmakeSessionResultOptionApplicationBuffer FindMatchmakeSessionResultOption = 1

	// FindMatchmakeSessionResultOptionMatchmakeParam populates m_MatchmakeParam in the results
	FindMatchmakeSessionResultOptionMatchmakeParam FindMatchmakeSessionResultOption = 2
)
