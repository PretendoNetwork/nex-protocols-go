package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NotificationType represents the type of notification being sent.
//
// Note: This is not a real type. This is a bespoke type made for our convenience.
type NotificationType uint32

// WriteTo writes the NotificationType to the given writable
func (nt NotificationType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(nt))
}

// ExtractFrom extracts the NotificationType value from the given readable
func (nt *NotificationType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*nt = NotificationType(value)
	return nil
}

// String returns a human-readable representation of the NotificationType.
func (nt NotificationType) String() string {
	switch nt {
	case NotificationTypeFriendPresenceUpdated3DS:
		return "FriendPresenceUpdated3DS"
	case NotificationTypeFriendFavoriteGameUpdated3DS:
		return "FriendFavoriteGameUpdated3DS"
	case NotificationTypeFriendCommentUpdated3DS:
		return "FriendCommentUpdated3DS"
	case NotificationTypeFriendMiiChanged3DS:
		return "FriendMiiChanged3DS"
	case NotificationTypeFriendProfileUpdated3DS:
		return "FriendProfileUpdated3DS"
	case NotificationTypeFriendshipCompleted3DS:
		return "FriendshipCompleted3DS"
	case NotificationTypeFriendshipRemoved3DS:
		return "FriendshipRemoved3DS"
	case NotificationTypeFriendSentInvitation3DS:
		return "FriendSentInvitation3DS"
	case NotificationTypeFriendOffline:
		return "FriendOffline"
	case NotificationTypeFriendMiiChangedWiiU:
		return "FriendMiiChangedWiiU"
	case NotificationTypeUnknown1MiiRelatedWiiU:
		return "Unknown1MiiRelatedWiiU"
	case NotificationTypeFriendPreferencesChangedWiiU:
		return "FriendPreferencesChangedWiiU"
	case NotificationTypeFriendStartedTitleWiiU:
		return "FriendStartedTitleWiiU"
	case NotificationTypeUnknown2FriendRequestRelatedWiiU:
		return "Unknown2FriendRequestRelatedWiiU"
	case NotificationTypeFriendshipCanceledWiiU:
		return "FriendshipCanceledWiiU"
	case NotificationTypeFriendRequestReceivedWiiU:
		return "FriendRequestReceivedWiiU"
	case NotificationTypeUnknown3FriendRequestRelatedWiiU:
		return "Unknown3FriendRequestRelatedWiiU"
	case NotificationTypeUnknown4BlacklistRelatedWiiU:
		return "Unknown4BlacklistRelatedWiiU"
	case NotificationTypeFriendRequestAcceptedWiiU:
		return "FriendRequestAcceptedWiiU"
	case NotificationTypeUnknown5BlacklistRelatedWiiU:
		return "Unknown5BlacklistRelatedWiiU"
	case NotificationTypeUnknown6BlacklistRelatedWiiU:
		return "Unknown6BlacklistRelatedWiiU"
	case NotificationTypeFriendStatusMessageChangedWiiU:
		return "FriendStatusMessageChangedWiiU"
	case NotificationTypeUnknown7WiiU:
		return "Unknown7WiiU"
	case NotificationTypeUnknown8FriendshipRelatedWiiU:
		return "Unknown8FriendshipRelatedWiiU"
	case NotificationTypeUnknown9PersistentNotificationsRelatedWiiU:
		return "Unknown9PersistentNotificationsRelatedWiiU"
	default:
		return fmt.Sprintf("NotificationType(%d)", int(nt))
	}
}

// * 3DS notifications
const (
	// NotificationTypeFriendPresenceUpdated3DS is delivered when a 3DS friend updates
	// their presence.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoPresence
	NotificationTypeFriendPresenceUpdated3DS NotificationType = 1

	// NotificationTypeFriendFavoriteGameUpdated3DS is delivered when a 3DS friend changes
	// their favorite game.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a GameKey.
	NotificationTypeFriendFavoriteGameUpdated3DS NotificationType = 2

	// NotificationTypeFriendCommentUpdated3DS is delivered when a 3DS friend changes their
	// comment.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendCommentUpdated3DS NotificationType = 3

	// NotificationTypeFriendMiiChanged3DS is delivered when a 3DS friend changes their Mii.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendMiiChanged3DS NotificationType = 5

	// NotificationTypeFriendProfileUpdated3DS is delivered when a 3DS friend updates their
	// profile.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventProfile.
	NotificationTypeFriendProfileUpdated3DS NotificationType = 6

	// NotificationTypeFriendshipCompleted3DS is delivered when you become connected friends
	// with another 3DS user.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendshipCompleted3DS NotificationType = 7

	// NotificationTypeFriendshipRemoved3DS is delivered when a 3DS friend delets your friend
	// card.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendshipRemoved3DS NotificationType = 8

	// NotificationTypeFriendSentInvitation3DS is delivered when a 3DS friend sends you a game
	// invitation.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendSentInvitation3DS NotificationType = 9
)

// * 3DS and Wii U notifications
const (
	// NotificationTypeFriendOffline is delivered when a friend goes offline.
	// This applies to both the 3DS and Wii U
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendOffline NotificationType = 10
)

// * Wii U notifications
const (
	// NotificationTypeFriendMiiChangedWiiU is delivered when a Wii U friend changes
	// their Mii.
	//
	// Sent using ProcessNintendoNotificationEvent 2.
	//
	// NintendoNotificationEvent.DataHolder is a NNAInfo.
	NotificationTypeFriendMiiChangedWiiU NotificationType = 21

	// NotificationTypeUnknown1MiiRelatedWiiU has an unknown use. Is related to Wii U
	// Miis.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown1MiiRelatedWiiU NotificationType = 22

	// NotificationTypeFriendPreferencesChangedWiiU is delivered when a Wii U friend
	// changes their preferences.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a PrincipalPreference.
	NotificationTypeFriendPreferencesChangedWiiU NotificationType = 23

	// NotificationTypeFriendStartedTitleWiiU is delivered when a Wii U friend opens
	// a title.
	//
	// Sent using ProcessNintendoNotificationEvent 2.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoPresenceV2.
	NotificationTypeFriendStartedTitleWiiU NotificationType = 24

	// NotificationTypeUnknown2FriendRequestRelatedWiiU has an unknown use. Is related to
	// Wii U friend requests.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown2FriendRequestRelatedWiiU NotificationType = 25

	// NotificationTypeFriendshipCanceledWiiU is delivered when a Wii U friend removes
	// you from their friends list, or when a pending incoming friend request is canceled.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendshipCanceledWiiU NotificationType = 26

	// NotificationTypeFriendRequestReceivedWiiU is delivered when a Wii U friend
	// request is received.
	//
	// Sent using ProcessNintendoNotificationEvent 2.
	//
	// NintendoNotificationEvent.DataHolder is a FriendRequest.
	NotificationTypeFriendRequestReceivedWiiU NotificationType = 27

	// NotificationTypeUnknown3FriendRequestRelatedWiiU has an unknown use. Is related to
	// Wii U friend requests.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown3FriendRequestRelatedWiiU NotificationType = 28

	// NotificationTypeUnknown4BlacklistRelatedWiiU has an unknown use. Is related to
	// the Wii U blacklist.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown4BlacklistRelatedWiiU NotificationType = 29

	// NotificationTypeFriendRequestAcceptedWiiU is delivered when you become friends with
	// another Wii U user.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a FriendInfo.
	NotificationTypeFriendRequestAcceptedWiiU NotificationType = 30

	// NotificationTypeUnknown5BlacklistRelatedWiiU has an unknown use. Is related to
	// the Wii U blacklist.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown5BlacklistRelatedWiiU NotificationType = 31

	// NotificationTypeUnknown6BlacklistRelatedWiiU has an unknown use. Is related to
	// the Wii U blacklist.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown6BlacklistRelatedWiiU NotificationType = 32

	// NotificationTypeFriendStatusMessageChangedWiiU is delivered when a Wii U friend
	// changes their status message.
	//
	// Sent using ProcessNintendoNotificationEvent 1.
	//
	// NintendoNotificationEvent.DataHolder is a NintendoNotificationEventGeneral.
	NotificationTypeFriendStatusMessageChangedWiiU NotificationType = 33

	// NotificationTypeUnknown7WiiU has an unknown use.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown7WiiU NotificationType = 34

	// NotificationTypeUnknown8FriendshipRelatedWiiU has an unknown use. Is related to
	// Wii U friendships.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder unknown.
	NotificationTypeUnknown8FriendshipRelatedWiiU NotificationType = 35

	// NotificationTypeUnknown9PersistentNotificationsRelatedWiiU has an unknown use. Seems
	// related to deleting persistent notifications.
	//
	// Send method unknown.
	//
	// NintendoNotificationEvent.DataHolder is a PersistentNotificationList.
	NotificationTypeUnknown9PersistentNotificationsRelatedWiiU NotificationType = 36
)
