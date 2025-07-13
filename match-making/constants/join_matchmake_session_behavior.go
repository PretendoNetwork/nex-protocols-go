package constants

// JoinMatchmakeSessionBehavior is used to indicate the behavior that joining a matchmake session will have
type JoinMatchmakeSessionBehavior uint8

// IsValid ensures the value of the JoinMatchmakeSessionBehavior is within
// the expected range
func (jmsb JoinMatchmakeSessionBehavior) IsValid() bool {
	return jmsb >= JoinMatchmakeSessionBehaviorJoinMyself && jmsb <= JoinMatchmakeSessionBehaviorImAlreadyJoined
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
