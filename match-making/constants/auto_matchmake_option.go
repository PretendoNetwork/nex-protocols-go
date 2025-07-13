package constants

// AutoMatchmakeOption has an unknown use.
type AutoMatchmakeOption uint32

// IsValid ensures the value of the AutoMatchmakeOption is within
// the expected range
func (amo AutoMatchmakeOption) IsValid() bool {
	return amo >= AutoMatchmakeOptionNone && amo <= AutoMatchmakeOptionUniqueGatheringByCodeword
}

const (
	// AutoMatchmakeOptionNone has an unknown use.
	AutoMatchmakeOptionNone AutoMatchmakeOption = iota

	// AutoMatchmakeOptionRecordLastGIDForParticipationCheck has an unknown use.
	//
	// Possibly related to gidForParticipationCheck?
	AutoMatchmakeOptionRecordLastGIDForParticipationCheck

	// AutoMatchmakeOptionUniqueGatheringByCodeword has an unknown use.
	//
	// Possibly related to m_Codeword?
	AutoMatchmakeOptionUniqueGatheringByCodeword
)
