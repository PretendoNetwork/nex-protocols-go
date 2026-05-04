package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubType exists solely to restrict the kinds of values that can be passed
// to SubType.Build()
type SubType uint32

// WriteTo writes the SubType to the given writable
func (st SubType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(st))
}

// ExtractFrom extracts the SubType value from the given readable
func (st *SubType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*st = SubType(value)
	return nil
}

// String returns a human-readable representation of SubType.
// This cannot differentiate between types of SubTypes
func (st SubType) String() string {
	return fmt.Sprintf("SubType(%d)", int(st))
}

// StringParticipationEvents returns a human-readable representation of the ParticipationEvents.
func (st SubType) StringParticipationEvents() string {
	pe := ParticipationEvents(st)
	switch pe {
	case ParticipationEventsParticipate:
		return "Participate"
	case ParticipationEventsCancelParticipation:
		return "CancelParticipation"
	case ParticipationEventsDisconnect:
		return "Disconnect"
	case ParticipationEventsEndParticipation:
		return "EndParticipation"
	default:
		return fmt.Sprintf("ParticipationEvents(%d)", int(pe))
	}
}

// StringSubscriptionEvents returns a human-readable representation of the SubscriptionEvents.
func (st SubType) StringSubscriptionEvents() string {
	se := SubscriptionEvents(st)
	switch se {
	case SubscriptionEventsEvent0:
		return "Event0"
	case SubscriptionEventsEvent1:
		return "Event1"
	case SubscriptionEventsEvent2:
		return "Event2"
	default:
		return fmt.Sprintf("SubscriptionEvents(%d)", int(se))
	}
}
