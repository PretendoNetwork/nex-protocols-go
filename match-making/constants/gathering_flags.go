package constants

// GatheringFlags indicates the flags set on a gathering
type GatheringFlags uint16

const (
	// GatheringFlagNone means that no flags are set
	GatheringFlagNone GatheringFlags = 0x0

	// GatheringFlagPersistentGathering indicates that a gathering is persistent.
	//
	// Gatherings with this flag don't get deleted on the server when the user logs
	// out or when the gathering state is set to "Finished". The gathering MUST be
	// deleted manually by the owner.
	//
	// Seen in Mario Kart 7 communities.
	GatheringFlagPersistentGathering GatheringFlags = 0x1

	// GatheringFlagMigrateOwner determines whether or not to pick a new owner when
	// the current owner logs out. If set and the owner disconnects, a new owner is
	// picked from the remaining participants. If a new owner can't be set, or this
	// flag is not set, then the gathering is deleted when the owner disconnects.
	//
	// Only applies to non-persistent gatherings.
	GatheringFlagMigrateOwner GatheringFlags = 0x10

	// GatheringFlagNoPersistentParticipation determines whether or not a participant
	// will have their participation status removed from a persistent gathering when
	// they log out. If set, the users participation status is removed after disconnecting.
	GatheringFlagNoPersistentParticipation GatheringFlags = 0x40

	// GatheringFlagAllowNoParticipant determines whether or not a persistent gathering
	// will be deleted when it has 0 participants. Normally a persistent gathering is
	// deleted when it has 0 participants. If set, the persistent gathering will NOT be
	// deleted when it has 0 participants.
	//
	// Note: This does NOT mean ACTIVE players. Players may be participants of a persistent
	// gathering while not actively playing in the session or even online.
	GatheringFlagAllowNoParticipant GatheringFlags = 0x80

	// GatheringFlagChangeOwnerByOtherHost determines whether or not ownership of a gathering
	// can be taken by another user. If set, calling UpdateSessionHost and UpdateGatheringOwnership
	// will allow ownership to be transfered. If set and ownership changes, a OwnershipChangeEvent
	// notification is sent to all participants.
	GatheringFlagChangeOwnerByOtherHost GatheringFlags = 0x200

	// GatheringFlagNotifyParticipationEventsToAllParticipants determines whether or
	// not Participate, Disconnect, and EndParticipation ParticipationEvents notifications
	// are sent to players in the session. If neither this flag nor
	// GatheringFlagNotifyParticipationEventsToAllParticipantsReproducibly are set then
	// only the owner of the session gets these notifications.
	GatheringFlagNotifyParticipationEventsToAllParticipants GatheringFlags = 0x400

	// GatheringFlagNotifyParticipationEventsToAllParticipantsReproducibly determines
	// whether or not Participate, Disconnect, and EndParticipation ParticipationEvents
	// notifications are sent to players in the session. If neither this flag nor
	// GatheringFlagNotifyParticipationEventsToAllParticipants are set then only the
	// owner of the session gets these notifications.
	//
	// The exact details on how this differs from GatheringFlagNotifyParticipationEventsToAllParticipants
	// is not known. It seems to seend a duplicate Participate notification event to
	// all existing players of the session (but not the new player)?
	GatheringFlagNotifyParticipationEventsToAllParticipantsReproducibly GatheringFlags = 0x800
)
