package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AutoMatchmakeOption has an unknown use.
type AutoMatchmakeOption uint32

// WriteTo writes the AutoMatchmakeOption to the given writable
func (amo AutoMatchmakeOption) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(amo))
}

// ExtractFrom extracts the AutoMatchmakeOption value from the given readable
func (amo *AutoMatchmakeOption) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*amo = AutoMatchmakeOption(value)
	return nil
}

// String returns a human-readable representation of the AutoMatchmakeOption.
func (amo AutoMatchmakeOption) String() string {
	switch amo {
	case AutoMatchmakeOptionNone:
		return "None"
	case AutoMatchmakeOptionRecordLastGIDForParticipationCheck:
		return "RecordLastGIDForParticipationCheck"
	case AutoMatchmakeOptionUniqueGatheringByCodeword:
		return "UniqueGatheringByCodeword"
	default:
		return fmt.Sprintf("AutoMatchmakeOption(%d)", int(amo))
	}
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
