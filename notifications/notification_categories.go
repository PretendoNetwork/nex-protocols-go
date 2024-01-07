// Package protocol implements the Notifications protocol
package protocol

type notificationCategories struct {
	Participation                *types.PrimitiveU32
	OwnershipChanged             *types.PrimitiveU32
	RequestJoinGathering         *types.PrimitiveU32 // * This is what these mean in WiiU Chat, unclear if this is the real use
	EndGathering                 *types.PrimitiveU32 // * This is what these mean in WiiU Chat, unclear if this is the real use
	GatheringUnregistered        *types.PrimitiveU32
	HostChanged                  *types.PrimitiveU32
	ServiceItemRequestCompleted  *types.PrimitiveU32
	MatchmakeRefereeRoundStarted *types.PrimitiveU32
	SystemPasswordChanged        *types.PrimitiveU32
	SystemPasswordCleared        *types.PrimitiveU32
	SwitchGathering              *types.PrimitiveU32
}

// NotificationCategories is a list of all the categories a notification can be in the Notifications protocol
//
// Not all of these are categories, some are stand-alone types
// This is a design choice made by NEX, not us
var NotificationCategories = notificationCategories{
	Participation:                3,
	OwnershipChanged:             4,
	RequestJoinGathering:         101, // * This is what these mean in WiiU Chat, unclear if this is the real use
	EndGathering:                 102, // * This is what these mean in WiiU Chat, unclear if this is the real use
	GatheringUnregistered:        109,
	HostChanged:                  110,
	ServiceItemRequestCompleted:  115,
	MatchmakeRefereeRoundStarted: 116,
	SystemPasswordChanged:        120,
	SystemPasswordCleared:        121,
	SwitchGathering:              122,
}
