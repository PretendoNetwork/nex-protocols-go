package constants

// NotificationEvents represents the main category of a notification type.
//
// Notifications may also include an optional subtype.
//
// Some games may make modifications to the meanings of the category parameters,
// or add additional parameters. The documentation here describes the default
// behavior, or a best guess.
type NotificationEvents uint32

// subType exists solely to restrict the kinds of values that can be passed
// to NotificationEvents.Build()
type subType uint32

func (ne NotificationEvents) Build(subtype ...subType) NotificationEvents {
	category := ne * 1000

	if len(subtype) == 0 {
		return category
	}

	return category + NotificationEvents(subtype[0])
}

const (
	// NotificationEventsSessionLaunched is delivered to everyone in a gathering
	// when MatchMaking::LaunchSession is fired.
	//
	// Parameters unknown.
	NotificationEventsSessionLaunched NotificationEvents = 2

	// NotificationEventsParticipationEvent is delivered when an event relating to
	// a users participation in a gathering has been updated. Who this event is
	// delivered to is determined by the gatherings flags.
	//
	// This category contains subtypes, see ParticipationEvents for details.
	NotificationEventsParticipationEvent NotificationEvents = 3

	// NotificationEventsOwnershipChangeEvent is delivered to everyone in a gathering
	// when the owner of a gathering has been changed.
	//
	// The parameters are:
	//
	// - m_pidSource: The previous owner
	// - m_uiParam1: The gathering ID
	// - m_uiParam2: The new owner
	NotificationEventsOwnershipChangeEvent NotificationEvents = 4

	// NotificationEventsGameNotification1 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification1 NotificationEvents = 101

	// NotificationEventsGameNotification2 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification2 NotificationEvents = 102

	// NotificationEventsGameNotification3 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification3 NotificationEvents = 103

	// NotificationEventsGameNotification4 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification4 NotificationEvents = 104

	// NotificationEventsGameNotification5 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification5 NotificationEvents = 105

	// NotificationEventsGameNotification6 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification6 NotificationEvents = 106

	// NotificationEventsGameNotification7 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification7 NotificationEvents = 107

	// NotificationEventsGameNotification8 is reserved for game-specific use.
	//
	// Note: Name is a guess, as we don't know the real name. Guess is based
	// on the fact that NotificationEventsGameNotificationLogout seems to only
	// be delivered when MatchmakeExtension::UpdateNotificationData has been
	// called by the user, and MatchmakeExtension::UpdateNotificationData sends
	// notifications of this type, and the  event has "GameNotification" in the name.
	NotificationEventsGameNotification8 NotificationEvents = 108

	// NotificationEventsGatheringUnregistered is delivered to everyone in a gathering
	// when the gathering has been unregistered.
	//
	// The parameters are:
	//
	// - m_pidSource: The gathering owner
	// - m_uiParam1: The gathering ID
	NotificationEventsGatheringUnregistered NotificationEvents = 109

	// NotificationEventsHostChangeEvent is delivered to the old host when a new host
	// has been selected.
	//
	// The parameters are:
	//
	// - m_pidSource: The new host?
	// - m_uiParam1: The gathering ID
	NotificationEventsHostChangeEvent NotificationEvents = 110

	// NotificationEventsGameNotificationLogout is delivered to the users friends when
	// the user disconnects, but seemingly only when MatchmakeExtension::UpdateNotificationData
	// has been called in the past by the disconnecting user?
	//
	// Parameters unknown, though likely have to do with the data sent in
	// MatchmakeExtension::UpdateNotificationData?
	NotificationEventsGameNotificationLogout NotificationEvents = 111

	// NotificationEventsSubscriptionEvent is delivered when an event relating to
	// the Subscription protocol happens.
	//
	// Details unknown.
	//
	// This category contains subtypes, see SubscriptionEvents for details.
	NotificationEventsSubscriptionEvent NotificationEvents = 112

	// NotificationEventsGameServerMaintenance is delivered to presumably all connected
	// clients to alert them of game server maintenance starting?
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationEventsGameServerMaintenance NotificationEvents = 113

	// NotificationEventsMaintenanceAnnouncement is delivered to presumably all connected
	// clients to alert them of future game server maintenance?
	//
	// Details unknown.
	//
	// Parameters unknown, though one of the parameters is likely a timestamp for when
	// maintenance begins?
	NotificationEventsMaintenanceAnnouncement NotificationEvents = 114

	// NotificationEventsServiceItemRequestCompleted is delivered to the caller of ServiceItem
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
	NotificationEventsServiceItemRequestCompleted NotificationEvents = 115

	// NotificationEventsRoundStarted is delivered to everyone in a gathering when a
	// when a MatchmakeReferee round is started.
	//
	// The parameters are:
	//
	// - m_pidSource: The caller of MatchmakeReferee::StartRound
	// - m_uiParam1: The round ID
	NotificationEventsRoundStarted NotificationEvents = 116

	// NotificationEventsFirstRoundReportReceived is delivered to everyone in a
	// gathering when a player calls MatchmakeReferee::EndRound for the first time. No
	// subsequent calls from other players in the round trigger this notification. No
	// calls to MatchmakeReferee::EndRoundWithoutReport trigger this notification.
	//
	// The parameters are:
	//
	// - m_pidSource: The caller of MatchmakeReferee::EndRound
	// - m_uiParam1: The round ID
	NotificationEventsFirstRoundReportReceived NotificationEvents = 117

	// NotificationEventsRoundSummarized is delivered to everyone in a gathering when
	// the rounds results are summarized. Round results seem to be summarized in 2 ways:
	//
	// 1. All players have called MatchmakeReferee::EndRound
	// 2. Some time (a couple minutes?) has past since the first MatchmakeReferee::EndRound call
	//
	// The parameters are:
	//
	// - m_uiParam1: The round ID
	NotificationEventsRoundSummarized NotificationEvents = 118

	// NotificationEventsMatchmakeSystemConfigurationNotification has an unknown use.
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationEventsMatchmakeSystemConfigurationNotification NotificationEvents = 119

	// NotificationEventsMatchmakeSessionSystemPasswordSet is delivered to everyone in
	// a gathering when the session system password has been updated.
	//
	// The parameters are:
	//
	// - m_pidSource: The updater
	// - m_strParam: The new password
	NotificationEventsMatchmakeSessionSystemPasswordSet NotificationEvents = 120

	// NotificationEventsMatchmakeSessionSystemPasswordClear is delivered to everyone in
	// a gathering when the session system password has been cleared.
	//
	// The parameters are:
	//
	// - m_pidSource: The clearer
	NotificationEventsMatchmakeSessionSystemPasswordClear NotificationEvents = 121

	// NotificationEventsAddedToGathering is delivered to a user when they are being added
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
	NotificationEventsAddedToGathering NotificationEvents = 122

	// NotificationEventsUserStatusUpdatedEvent has an unknown use. Possibly related
	// to Subscriber::UpdateUserStatus?
	//
	// Details unknown.
	//
	// Parameters unknown.
	NotificationEventsUserStatusUpdatedEvent NotificationEvents = 128
)
