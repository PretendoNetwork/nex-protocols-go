package constants

const (
	// TicketGrantingMaxUsernameLength is the max length a username
	// can be. If a username is larger than this value, the RMC
	// connection closes with no error response.
	//
	// Note: This is not a real type. This is a bespoke type created
	// for our convenience.
	TicketGrantingMaxUsernameLength int = 969 // * This is based on the friends server, which seems to only allow usernames of 970 bytes (969 + 1 null byte)

	// TicketGrantingMaxTokenLength is the max length a token
	// can be. If a token is larger than this value, the RMC
	// connection closes with no error response.
	//
	// Note: This is not a real type. This is a bespoke type created
	// for our convenience.
	TicketGrantingMaxTokenLength int = 918 // * This is based on the friends server, which seems to only allow usernames of 919 bytes (918 + 1 null byte)
)
