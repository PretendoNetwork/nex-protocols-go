package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UpdateMode is used by RankingScoreData.UpdateMode to control if worse scores should be discarded or not.
type UpdateMode uint8

// WriteTo writes the UpdateMode to the given writable
func (um UpdateMode) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(um))
}

// ExtractFrom extracts the UpdateMode value from the given readable
func (um *UpdateMode) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*um = UpdateMode(value)
	return nil
}

// String returns a human-readable representation of the UpdateMode.
func (um UpdateMode) String() string {
	switch um {
	case UpdateModeNormal:
		return "Normal"
	case UpdateModeDeleteOld:
		return "DeleteOld"
	default:
		return fmt.Sprintf("UpdateMode(%d)", int(um))
	}
}

const (
	// UpdateModeNormal will only commit the updated score if it's better than the old score in that category.
	// "better" is determined by the ranking order of the category.
	UpdateModeNormal UpdateMode = iota

	// UpdateModeDeleteOld will always overwrite any old score, even if the new one is worse.
	UpdateModeDeleteOld
)
