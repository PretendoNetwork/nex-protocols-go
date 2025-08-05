package constants

// GatheringState indicates the state of a gathering.
//
// Note: We do not know the real names for any of these, this is
// all guess work. The values of these states seem to imply that
// they are bitwise flags, but they seem to be treated as enum
// values in practice?
type GatheringState uint32

// IsValid ensures the value of the GatheringState is within
// the expected range
func (pp GatheringState) IsValid() bool {
	// * Kinda jank but whatever, screw it. Once we know all the
	// * real values of this enum we can do it like the others
	switch pp {
	case GatheringStateClosed:
	case GatheringStateStarted:
	case GatheringStateFinished:
		return true
	}

	return false
}

const (
	// GatheringStateClosed indicates that a gathering is closed, and no new participants may join.
	GatheringStateClosed GatheringState = 1

	// GatheringStateStarted indicates that the gatherings session has begun.
	GatheringStateStarted GatheringState = 2

	// GatheringStateFinished indicates that the gatherings session has finished.
	// Manually setting this state seems to delete gatherings which are non-persistent.
	GatheringStateFinished GatheringState = 4
)
