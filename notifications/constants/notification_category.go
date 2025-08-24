package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// NotificationCategory represents the main category of a notification type.
//
// Notifications may also include an optional subtype.
//
// Some games may make modifications to the meanings of the category parameters,
// or add additional parameters. The documentation here describes the default
// behavior, or a best guess.
//
// Note: The official name of this type is NotificationEvents. This has been changed
// to a slightly more intuitive name, and NotificationEvents has been repurposed to
// represent the fully built notification type with it's sub-type.
type NotificationCategory uint32

// WriteTo writes the NotificationCategory to the given writable
func (nc NotificationCategory) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(nc))
}

// ExtractFrom extracts the NotificationCategory value from the given readable
func (nc *NotificationCategory) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*nc = NotificationCategory(value)
	return nil
}

// Build creates the final notification type ID used in NotificationEvent.m_uiType.
//
// Takes an optional subtype. Only the first subtype defined is used.
func (nc NotificationCategory) Build(subtype ...SubType) NotificationEvents {
	category := NotificationEvents(nc * 1000)

	if len(subtype) == 0 {
		return category
	}

	return category + NotificationEvents(subtype[0])
}

// ToSigned converts a NotificationCategory to a NotificationCategorySigned
func (nc NotificationCategory) ToSigned() NotificationCategorySigned {
	return NotificationCategorySigned(nc)
}

const (
	// NotificationCategorySessionLaunched is delivered to everyone in a gathering
	// when MatchMaking::LaunchSession is fired.
	//
	// Parameters unknown.
	NotificationCategorySessionLaunched NotificationCategory = 2

	// NotificationCategoryParticipationEvent is delivered when an event relating to
	// a users participation in a gathering has been updated. Who this event is
	// delivered to is determined by the gatherings flags.
	//
	// This category contains subtypes, see ParticipationEvents for details.
	NotificationCategoryParticipationEvent NotificationCategory = 3

	// NotificationCategoryOwnershipChangeEvent is delivered to everyone in a gathering
	// when the owner of a gathering has been changed.
	//
	// The parameters are:
	//
	// - m_pidSource: The previous owner
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The new owner
	NotificationCategoryOwnershipChangeEvent NotificationCategory = 4

	// NotificationCategoryGameNotification1 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification1 NotificationCategory = 101

	// NotificationCategoryGameNotification2 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification2 NotificationCategory = 102

	// NotificationCategoryGameNotification3 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification3 NotificationCategory = 103

	// NotificationCategoryGameNotification4 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification4 NotificationCategory = 104

	// NotificationCategoryGameNotification5 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification5 NotificationCategory = 105

	// NotificationCategoryGameNotification6 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification6 NotificationCategory = 106

	// NotificationCategoryGameNotification7 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification7 NotificationCategory = 107

	// NotificationCategoryGameNotification8 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationCategoryGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationCategoryGameNotification8 NotificationCategory = 108

	// NotificationCategoryGatheringUnregistered is delivered to everyone in a gathering
	// when the gathering has been unregistered.
	//
	// The parameters are:
	//
	// - m_pidSource: The gathering owner
	// - m_uiParam1: The gathering ID
	NotificationCategoryGatheringUnregistered NotificationCategory = 109

	// NotificationCategoryHostChangeEvent is delivered to the old host when a new host
	// has been selected.
	//
	// The parameters are:
	//
	// - m_pidSource: The new host?
	// - m_uiParam1: The gathering ID
	NotificationCategoryHostChangeEvent NotificationCategory = 110

	// NotificationCategoryGameNotificationLogout is delivered to the users friends when
	// the user disconnects, but seemingly only when MatchmakeExtension::UpdateNotificationData
	// has been called in the past by the disconnecting user?
	//
	// Parameters unknown, though likely have to do with the data sent in
	// MatchmakeExtension::UpdateNotificationData?
	NotificationCategoryGameNotificationLogout NotificationCategory = 111

	// NotificationCategorySubscriptionEvent is delivered when an event relating to
	// the Subscription protocol happens.
	//
	// Details unknown.
	//
	// This category contains subtypes, see SubscriptionEvents for details.
	NotificationCategorySubscriptionEvent NotificationCategory = 112

	// NotificationCategoryGameServerMaintenance is delivered to presumably all connected
	// clients to alert them of game server maintenance starting?
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationCategoryGameServerMaintenance NotificationCategory = 113

	// NotificationCategoryMaintenanceAnnouncement is delivered to presumably all connected
	// clients to alert them of future game server maintenance?
	//
	// Details unknown.
	//
	// Parameters unknown, though one of the parameters is likely a timestamp for when
	// maintenance begins?
	NotificationCategoryMaintenanceAnnouncement NotificationCategory = 114

	// NotificationCategoryServiceItemRequestCompleted is delivered to the caller of ServiceItem
	// protocol "Request" events?
	//
	// Details unknown.
	//
	// The parameters are:
	//
	// - m_pidSource: The calling user?
	// - m_uiParam1: The request ID?
	// - m_uiParam2: Unknown. Always 1?
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based on the usage context.
	NotificationCategoryServiceItemRequestCompleted NotificationCategory = 115

	// NotificationCategoryRoundStarted is delivered to everyone in a gathering when a
	// when a MatchmakeReferee round is started.
	//
	// The parameters are:
	//
	// - m_pidSource: The caller of MatchmakeReferee::StartRound
	// - m_uiParam1: The round ID
	NotificationCategoryRoundStarted NotificationCategory = 116

	// NotificationCategoryFirstRoundReportReceived is delivered to everyone in a
	// gathering when a player calls MatchmakeReferee::EndRound for the first time. No
	// subsequent calls from other players in the round trigger this notification. No
	// calls to MatchmakeReferee::EndRoundWithoutReport trigger this notification.
	//
	// The parameters are:
	//
	// - m_pidSource: The caller of MatchmakeReferee::EndRound
	// - m_uiParam1: The round ID
	NotificationCategoryFirstRoundReportReceived NotificationCategory = 117

	// NotificationCategoryRoundSummarized is delivered to everyone in a gathering when
	// the rounds results are summarized. Round results seem to be summarized in 2 ways:
	//
	// 1. All players have called MatchmakeReferee::EndRound
	// 2. Some time (a couple minutes?) has past since the first MatchmakeReferee::EndRound call
	//
	// The parameters are:
	//
	// - m_uiParam1: The round ID
	NotificationCategoryRoundSummarized NotificationCategory = 118

	// NotificationCategoryMatchmakeSystemConfigurationNotification has an unknown use.
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationCategoryMatchmakeSystemConfigurationNotification NotificationCategory = 119

	// NotificationCategoryMatchmakeSessionSystemPasswordSet is delivered to everyone in
	// a gathering when the session system password has been updated.
	//
	// The parameters are:
	//
	// - m_pidSource: The updater
	// - m_strParam: The new password
	NotificationCategoryMatchmakeSessionSystemPasswordSet NotificationCategory = 120

	// NotificationCategoryMatchmakeSessionSystemPasswordClear is delivered to everyone in
	// a gathering when the session system password has been cleared.
	//
	// The parameters are:
	//
	// - m_pidSource: The clearer
	NotificationCategoryMatchmakeSessionSystemPasswordClear NotificationCategory = 121

	// NotificationCategoryAddedToGathering is delivered to a user when they are being added
	// as a participant of a gathering. Users may be participants in more than one gatering
	// at a time, such as with persistent gatherings. This does not *move* a user *from* an
	// existing gathering, only *adds* to the new one.
	//
	// The parameters are:
	//
	// - m_pidSource: The player who called the method
	// - m_uiParam1: The gathering ID the player is being added to
	// - m_uiParam2: The PID of the player being added
	// - m_strParam: The join message set when adding the new player
	NotificationCategoryAddedToGathering NotificationCategory = 122

	// NotificationCategoryUserStatusUpdatedEvent has an unknown use. Possibly related
	// to Subscriber::UpdateUserStatus?
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationCategoryUserStatusUpdatedEvent NotificationCategory = 128

	// NotificationCategoryEagleAddress is delivered to all clients of a session when the Eagle
	// server is used and is ready to be connected to. For details of the Eagle protocol see:
	// https://nintendo-wiki.pretendo.network/docs/switch/eagle/
	//
	// The parameters are:
	//
	// - m_pidSource: The secure server (257049437023956657 for the official Switch servers)
	// - m_uiParam1: The gathering ID
	// - m_mapParam: A map with 2 keys, "url" (the address of the Eagle server) and "token" (connection token)
	NotificationCategoryEagleAddress NotificationCategory = 200
)
