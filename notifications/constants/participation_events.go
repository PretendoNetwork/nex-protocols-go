package constants

// ParticipationEvents represents the various subtypes of NotificationEventsParticipationEvent
type ParticipationEvents = subType

const (
	// ParticipationParticipate is delivered when a new player has joined
	// the gathering.
	//
	// The parameters are:
	//
	// - m_pidSource: The user who initiated the join
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The joining player PID
	// - m_strParam: The join message
	// - m_uiParam3: The number of participants
	ParticipationParticipate ParticipationEvents = 1

	// ParticipationEventsCancelParticipation is delivered when a player
	// cancels their participation through MatchMake::CancelParticipation.
	//
	// The parameters are:
	//
	// - m_pidSource: The cancelling player
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The cancelling player PID
	// - m_strParam: The message sent in the RMC call
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the structure of ParticipationEventsEndParticipation.
	ParticipationEventsCancelParticipation ParticipationEvents = 2

	// ParticipationEventsDisconnect is delivered when a player disconnects
	// without first properly leaving the gathering, such as a lost connection.
	//
	// The parameters are:
	//
	// - m_pidSource: The disconnected player
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The disconnected player PID
	ParticipationEventsDisconnect ParticipationEvents = 7

	// ParticipationEventsEndParticipation is delivered when a player has decided
	// to leave the gathering gracefully, such as through MatchMakingExt::EndParticipation
	//
	// The parameters are:
	//
	// - m_pidSource: The leaving player
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The leaving player PID
	// - m_strParam: The message sent in the RMC call
	ParticipationEventsEndParticipation ParticipationEvents = 8
)
