package constants

// AutoMatchmakeOption has an unknown use.
type AutoMatchmakeOption uint32

func (amo AutoMatchmakeOption) IsValid() bool {
	checkValue := uint32(amo)
	return checkValue >= uint32(AutoMatchmakeOptionNone) && checkValue <= uint32(AutoMatchmakeOptionUniqueGatheringByCodeword)
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
