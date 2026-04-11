package constants

// * Miscellaneous constants

const (
	// MaxMiiNameSize is the maximum length of a Mii nickname, minus the null terminator
	MaxMiiNameSize int = 10

	// MaxApplicationDataSize is the maximum length of a NintendoPresence.m_applicationArg buffer
	MaxApplicationDataSize int = 20

	// MaxCommentSize is the maximum length of a comment message, minus the null terminator
	MaxCommentSize int = 16

	// MaxFriends is the maximum number of friends a user can have
	MaxFriends int = 100

	// CFLStoreDataSize is the size of Mii data stored in the CFLStoreData format
	CFLStoreDataSize int = 96

	// MaxSerialNumberSize is the maximum length of a console serial number, minus the null terminator
	MaxSerialNumberSize int = 15

	// MaxMACAddressSize is the maximum length of a console MAC address, minus the null terminator
	MaxMACAddressSize int = 17

	// MaxGameModeDescriptionSize is the maximum length of a `m_gameModeDescription`, minus the null terminator
	MaxGameModeDescriptionSize int = 127
)
