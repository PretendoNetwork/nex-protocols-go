package constants

// * Miscellaneous constants

const (
	// MaxAccountIDSize is the maximum length of a network account username, minus the null terminator
	MaxAccountIDSize int = 16

	// MaxMiiNameSize is the maximum length of a Mii nickname, minus the null terminator
	MaxMiiNameSize int = 10

	// MaxApplicationDataSize is the maximum length of a NintendoPresenceV2.ApplicationData buffer
	MaxApplicationDataSize int = 20

	// MaxCommentSize is the maximum length of a Comment message, minus the null terminator
	MaxCommentSize int = 16

	// MaxFriends is the maximum number of friends a user can have
	MaxFriends int = 100

	// FFLStoreDataSize is the size of Mii data stored in the FFLStoreData format
	FFLStoreDataSize int = 96

	// MaxFriendRequestMessageSize is the maximum length of a friend request message, minus the null terminator
	MaxFriendRequestMessageSize int = 63
)
