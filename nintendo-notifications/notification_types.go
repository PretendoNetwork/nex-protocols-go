// Package protocol implements the Nintendo Notfications protocol
package protocol

type notificationTypes struct {
	FriendPresenceUpdated3DS               *types.PrimitiveU32
	FriendFavoriteGameUpdated3DS           *types.PrimitiveU32
	FriendCommentUpdated3DS                *types.PrimitiveU32
	FriendMiiChanged3DS                    *types.PrimitiveU32
	FriendshipCompleted3DS                 *types.PrimitiveU32
	FriendOffline                          *types.PrimitiveU32
	FriendMiiChanged                       *types.PrimitiveU32
	Unknown1MiiRelated                     *types.PrimitiveU32
	FriendPreferencesChanged               *types.PrimitiveU32
	FriendStartedTitle                     *types.PrimitiveU32
	Unknown2FriendRequestRelated           *types.PrimitiveU32
	FriendRemoved                          *types.PrimitiveU32 // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestCanceled                  *types.PrimitiveU32 // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestReceived                  *types.PrimitiveU32
	Unknown3FriendRequestRelated           *types.PrimitiveU32
	Unknown4BlacklistRelated               *types.PrimitiveU32
	FriendRequestAccepted                  *types.PrimitiveU32
	Unknown5BlacklistRelated               *types.PrimitiveU32
	Unknown6BlacklistRelated               *types.PrimitiveU32
	FriendStatusMessageChanged             *types.PrimitiveU32
	Unknown7                               *types.PrimitiveU32
	Unknown8FriendshipRelated              *types.PrimitiveU32
	Unknown9PersistentNotificationsRelated *types.PrimitiveU32
}

// NotificationTypes is an enum of all the types a notification can be in the NintendoNotifications protocol
var NotificationTypes = notificationTypes{
	FriendPresenceUpdated3DS:               1,
	FriendFavoriteGameUpdated3DS:           2,
	FriendCommentUpdated3DS:                3,
	FriendMiiChanged3DS:                    5,
	FriendshipCompleted3DS:                 7,
	FriendOffline:                          10,
	FriendMiiChanged:                       21,
	Unknown1MiiRelated:                     22,
	FriendPreferencesChanged:               23,
	FriendStartedTitle:                     24,
	Unknown2FriendRequestRelated:           25,
	FriendRemoved:                          26, // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestCanceled:                  26, // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestReceived:                  27,
	Unknown3FriendRequestRelated:           28,
	Unknown4BlacklistRelated:               29,
	FriendRequestAccepted:                  30,
	Unknown5BlacklistRelated:               31,
	Unknown6BlacklistRelated:               32,
	FriendStatusMessageChanged:             33,
	Unknown7:                               34,
	Unknown8FriendshipRelated:              35,
	Unknown9PersistentNotificationsRelated: 36,
}
