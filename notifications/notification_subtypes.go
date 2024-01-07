// Package protocol implements the Notifications protocol
package protocol

type notificationSubTypes struct {
	Participation                notificationParticipationSubTypes
	OwnershipChanged             notificationsOwnershipChangedSubTypes
	RequestJoinGathering         notificationsRequestJoinGatheringSubTypes
	EndGathering                 notificationsEndGatheringSubTypes
	GatheringUnregistered        notificationsGatheringUnregisteredSubTypes
	HostChanged                  notificationsHostChangedSubTypes
	ServiceItemRequestCompleted  notificationsServiceItemRequestCompletedSubTypes
	MatchmakeRefereeRoundStarted notificationsMatchmakeRefereeRoundStartedSubTypes
	SystemPasswordChanged        notificationsSystemPasswordChangedSubTypes
	SystemPasswordCleared        notificationsSystemPasswordClearedSubTypes
	SwitchGathering              notificationsSwitchGatheringSubTypes
}

type notificationParticipationSubTypes struct {
	None           *types.PrimitiveU32
	NewParticipant *types.PrimitiveU32
	Cancelled      *types.PrimitiveU32
	Disconnected   *types.PrimitiveU32
	Ended          *types.PrimitiveU32
}

type notificationsOwnershipChangedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsRequestJoinGatheringSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsEndGatheringSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsGatheringUnregisteredSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsHostChangedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsServiceItemRequestCompletedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsMatchmakeRefereeRoundStartedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsSystemPasswordChangedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsSystemPasswordClearedSubTypes struct {
	None *types.PrimitiveU32
}

type notificationsSwitchGatheringSubTypes struct {
	None *types.PrimitiveU32
}

// NotificationSubTypes is a list of all the subtypes a notification can be in a given category of the Notifications protocol
//
// Not all categories have subtypes
var NotificationSubTypes = notificationSubTypes{
	Participation: notificationParticipationSubTypes{
		None:           0,
		NewParticipant: 1,
		Cancelled:      2,
		Disconnected:   7,
		Ended:          8,
	},
	OwnershipChanged: notificationsOwnershipChangedSubTypes{
		None: 0,
	},
	RequestJoinGathering: notificationsRequestJoinGatheringSubTypes{
		None: 0,
	},
	EndGathering: notificationsEndGatheringSubTypes{
		None: 0,
	},
	GatheringUnregistered: notificationsGatheringUnregisteredSubTypes{
		None: 0,
	},
	HostChanged: notificationsHostChangedSubTypes{
		None: 0,
	},
	ServiceItemRequestCompleted: notificationsServiceItemRequestCompletedSubTypes{
		None: 0,
	},
	MatchmakeRefereeRoundStarted: notificationsMatchmakeRefereeRoundStartedSubTypes{
		None: 0,
	},
	SystemPasswordChanged: notificationsSystemPasswordChangedSubTypes{
		None: 0,
	},
	SystemPasswordCleared: notificationsSystemPasswordClearedSubTypes{
		None: 0,
	},
	SwitchGathering: notificationsSwitchGatheringSubTypes{
		None: 0,
	},
}
