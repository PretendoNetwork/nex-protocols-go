package notifications

type notificationTypes struct {
	NewParticipant               uint32
	ParticipationCancelled       uint32
	ParticipantDisconnected      uint32
	ParticipationEnded           uint32
	OwnershipChanged             uint32
	GatheringUnregistered        uint32
	HostChanged                  uint32
	ServiceItemRequestCompleted  uint32
	MatchmakeRefereeRoundStarted uint32
	SystemPasswordChanged        uint32
	SystemPasswordCleared        uint32
	SwitchGathering              uint32
}

var NotificationTypes = notificationTypes{
	NewParticipant:               3001,
	ParticipationCancelled:       3002,
	ParticipantDisconnected:      3007,
	ParticipationEnded:           3008,
	OwnershipChanged:             4000,
	GatheringUnregistered:        109000,
	HostChanged:                  110000,
	ServiceItemRequestCompleted:  115000,
	MatchmakeRefereeRoundStarted: 116000,
	SystemPasswordChanged:        120000,
	SystemPasswordCleared:        121000,
	SwitchGathering:              122000,
}
