package constants

const (
	// MaxTopicContentSize is likely the maximum number of items that can be posted
	// to a topic?
	//
	// Note: This is entirely a guess, since constants like `MaxContentBinarySize`
	// seem to already exist to limit the size of the individual fields.
	MaxTopicContentSize uint32 = 100

	// MaxTimelineContentSize is likely the maximum number of items in a timeline?
	//
	// Note: This is entirely a guess, since it's not clear what a "timeline" is?
	// This is based off the assumption that MaxTopicContentSize means the maximum
	// number items in a topic. If that is true, then the same logic likely applies
	// here as well.
	MaxTimelineContentSize uint32 = 100

	// MaxFollowingSize is the maximum number of topics(?) that can be followed.
	//
	// Note: This is entirely a guess, since it's not clear if you can follow
	// both users and topics. So far it SEEMS like you only follow topics?
	MaxFollowingSize uint32 = 20

	// NumReservedTopics is the max value a topic ID can have?
	//
	// Note: This is entirely a guess. I assume "reserved topic" means that topics are registered
	// by the server, and this is the max value it can be?
	NumReservedTopics uint32 = 128

	// InvalidReservedTopicNum represents an invalid topic number.
	//
	// Note: This is entirely a guess, but is based on the logic in the `NumReservedTopics` guess.
	InvalidReservedTopicNum uint32 = 4294967295

	// MaxGetFollowerSize is the maximum number of followers that can be retrieved in `GetFollower`.
	//
	// Note: This is entirely a guess based on the name of the constant and RMC method.
	MaxGetFollowerSize uint32 = 1000

	// MaxContentMessageLen is the maximum length for `SubscriberPostContentParam.message`.
	MaxContentMessageLen uint32 = 140

	// MaxContentBinarySize is the maximum length for `SubscriberPostContentParam.binary`.
	MaxContentBinarySize uint32 = 256

	// MaxPostContentTopicSize is the maximum number of topics for `SubscriberPostContentParam.topic`.
	MaxPostContentTopicSize uint32 = 10

	// MaxGetContentParamsSize is the maximum number of parameters that can be sent to `GetContent`.
	//
	// Note: This is entirely a guess, and doesn't make much sense? `GetContent` only takes in a single value?
	// Is this maybe a client-side limit, where the client mades up to 3 individual `GetContent` requests? or
	// maybe it's a NEX version difference?
	MaxGetContentParamsSize uint32 = 3

	// StatusKeySize is the maximum number of `keys` list values. The values of `keys` may only
	// be up to this value as well (0-7).
	StatusKeySize uint8 = 8

	// MaxStatusBinarySize is the maximum length for user status qBuffers.
	MaxStatusBinarySize uint32 = 128

	// MaxGetStatusUserSize is the maximum number of parameters that can be sent to `GetUserStatuses`.
	MaxGetStatusUserSize uint32 = 100

	// DefaultRMCInterval is the default interval (in milliseconds?) which the client will poll for content?
	//
	// Note: Entirely a guess. I assume "Subscriber" means you "subscribe" the topics (and users?) and the client
	// polls for new data every so often.
	DefaultRMCInterval uint32 = 5000
)
