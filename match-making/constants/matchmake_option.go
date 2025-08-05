package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeOption has an unknown use.
type MatchmakeOption uint32

// WriteTo writes the MatchmakeOption to the given writable
func (mo MatchmakeOption) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(mo))
}

// ExtractFrom extracts the MatchmakeOption value from the given readable
func (mo *MatchmakeOption) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*mo = MatchmakeOption(value)
	if !mo.IsValid() {
		return fmt.Errorf("Value %d is out of range", *mo)
	}

	return nil
}

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
