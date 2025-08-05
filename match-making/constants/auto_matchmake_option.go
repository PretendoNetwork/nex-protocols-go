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
	if !amo.IsValid() {
		return fmt.Errorf("Value %d is out of range", *amo)
	}

	return nil
}

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
