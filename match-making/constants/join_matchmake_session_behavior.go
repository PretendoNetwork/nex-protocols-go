package constants

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// JoinMatchmakeSessionBehavior is used to indicate the behavior that joining a matchmake session will have
type JoinMatchmakeSessionBehavior uint8

// WriteTo writes the JoinMatchmakeSessionBehavior to the given writable
func (jmsb JoinMatchmakeSessionBehavior) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(jmsb))
}

// ExtractFrom extracts the JoinMatchmakeSessionBehavior value from the given readable
func (jmsb *JoinMatchmakeSessionBehavior) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*jmsb = JoinMatchmakeSessionBehavior(value)
	return nil
}

const (
	// JoinMatchmakeSessionBehaviorJoinMyself indicates that the caller wants to join the session.
	JoinMatchmakeSessionBehaviorJoinMyself JoinMatchmakeSessionBehavior = iota

	// JoinMatchmakeSessionBehaviorImAlreadyJoined indicates that the caller is already joined into the session.
	//
	// This can be useful in cases where the caller wants to add additional users to the session using additionalParticipants
	// without trying to join the session themselves multiple times.
	JoinMatchmakeSessionBehaviorImAlreadyJoined
)
