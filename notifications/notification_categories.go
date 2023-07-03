// Package notifications implements the Notifications NEX protocol
package notifications

type notificationCategories struct {
	Participation                uint32
	OwnershipChanged             uint32
	GatheringUnregistered        uint32
	HostChanged                  uint32
	ServiceItemRequestCompleted  uint32
	MatchmakeRefereeRoundStarted uint32
	SystemPasswordChanged        uint32
	SystemPasswordCleared        uint32
	SwitchGathering              uint32
}

// NotificationCategories is a list of all the categories a notification can be in the Notifications protocol
//
// Not all of these are categories, some are stand-alone types
// This is a design choice made by NEX, not us
var NotificationCategories = notificationCategories{
	Participation:                3,
	OwnershipChanged:             4,
	GatheringUnregistered:        109,
	HostChanged:                  110,
	ServiceItemRequestCompleted:  115,
	MatchmakeRefereeRoundStarted: 116,
	SystemPasswordChanged:        120,
	SystemPasswordCleared:        121,
	SwitchGathering:              122,
}
