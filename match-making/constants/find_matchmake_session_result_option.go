package constants

// FindMatchmakeSessionResultOption indicates how to populate the
// responses from FindMatchmakeSessionByParticipant
type FindMatchmakeSessionResultOption uint32

const (
	// FindMatchmakeSessionResultOptionNone indicates no options
	FindMatchmakeSessionResultOptionNone FindMatchmakeSessionResultOption = 0

	// FindMatchmakeSessionResultOptionApplicationBuffer populates m_ApplicationBuffer in the results
	FindMatchmakeSessionResultOptionApplicationBuffer FindMatchmakeSessionResultOption = 1

	// FindMatchmakeSessionResultOptionMatchmakeParam populates m_MatchmakeParam in the results
	FindMatchmakeSessionResultOptionMatchmakeParam FindMatchmakeSessionResultOption = 2
)
