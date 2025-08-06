package constants

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GatheringState indicates the state of a gathering.
//
// Note: We do not know the real names for any of these, this is
// all guess work. The values of these states seem to imply that
// they are bitwise flags, but they seem to be treated as enum
// values in practice?
type GatheringState uint32

// WriteTo writes the GatheringState to the given writable
func (gs GatheringState) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(gs))
}

// ExtractFrom extracts the GatheringState value from the given readable
func (gs *GatheringState) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*gs = GatheringState(value)
	return nil
}

const (
	// GatheringStateLocked indicates that a gathering is locked, and no new participants may join
	// Name assumed from MatchmakeSessionSearchCriteria.ExcludeLocked.
	GatheringStateLocked GatheringState = 1

	// GatheringStateStarted indicates that the gatherings session has begun.
	GatheringStateStarted GatheringState = 2

	// GatheringStateFinished indicates that the gatherings session has finished.
	// Manually setting this state seems to delete gatherings which are non-persistent.
	GatheringStateFinished GatheringState = 4
)
