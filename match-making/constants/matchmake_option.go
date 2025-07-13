package constants

// MatchmakeOption has an unknown use.
type MatchmakeOption uint32

// IsValid ensures the value of the MatchmakeOption is within
// the expected range
func (mo MatchmakeOption) IsValid() bool {
	return mo >= MatchmakeOptionNone && mo <= MatchmakeOptionReserved1
}

const (
	// MatchmakeOptionNone has an unknown use.
	MatchmakeOptionNone MatchmakeOption = iota

	// MatchmakeOptionRecordLastGIDForParticipationCheck has an unknown use.
	//
	// Possibly related to gidForParticipationCheck?
	MatchmakeOptionRecordLastGIDForParticipationCheck

	// MatchmakeOptionReserved1 has an unknown use.
	MatchmakeOptionReserved1
)
