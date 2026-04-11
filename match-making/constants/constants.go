package constants

// * Miscellaneous constants

const (
	// InvalidGatheringID represents an invalid/unset gathering ID
	InvalidGatheringID uint32 = 0

	// NumMatchmakeSessionAttributes is the max number of m_Attribs attributes a session may have
	NumMatchmakeSessionAttributes int = 6

	// MatchmakeSessionNearestNeighborAttributeIndex is the index into m_Attribs for finding the value
	// used for the SelectionMethodNearestNeighbor method
	MatchmakeSessionNearestNeighborAttributeIndex int = 1

	// MatchmakeSessionBroadenRangeAttributeIndex is the index into m_Attribs for finding the value
	// used for the SelectionMethodBroadenRange/SelectionMethodBroadenRangeWithProgressScore methods
	MatchmakeSessionBroadenRangeAttributeIndex int = 1

	// InvalidPersistentGatheringCode likely represents an invalid result from GatheringIdToPersistentGatheringCode?
	InvalidPersistentGatheringCode uint64 = 0

	// PersistentGatheringCreationMax is the max number of persistent gatherings a user can create
	PersistentGatheringCreationMax int = 4

	// PersistentGatheringParticipationMax is the max number of persistent gatherings a user can join
	PersistentGatheringParticipationMax int = 16

	// PersistentGatheringChatParticipantsMax is unknown. Maybe related to the "SendChat" methods of the NexMessagingClient?
	PersistentGatheringChatParticipantsMax int = 32

	// SetAttributeVectorSizeMax seems to be the max number of values a m_Attribs string may have? ("1,2,3,4,5" etc.)
	//
	// Note: This is a guess based on the "SetAttribute" and "VectorSize" parts. MatchmakeSessionSearchCriteria has a
	// SetAttribute method which sets the m_Attribs field. We also know that the m_Attribs values can contain multiple
	// numbers, possibly representing a "vector"
	SetAttributeVectorSizeMax int = 100

	// MaxProgressScore is the max value for progressScore
	MaxProgressScore uint8 = 100

	// UpdateProgressScoreMinimumIntervalTime is the minimum interval (in seconds?) for updating the sessions progressScore
	UpdateProgressScoreMinimumIntervalTime uint32 = 30

	// MaxMatchmakeSessionUserPasswordLength is the max user password length for a matchmake session
	MaxMatchmakeSessionUserPasswordLength int = 32

	// MatchmakeSessionSystemPasswordLength is the system password length for a matchmake session
	MatchmakeSessionSystemPasswordLength int = 16

	// MaxMatchmakeBrowseSize is the max number of gatherings that can be returned from the BrowseMatchmakeSession family of methods.
	//
	// This seems to also be the max input size for FindMatchmakeSessionByGatheringId?
	MaxMatchmakeBrowseSize int = 100

	// MaxPrincipalIDSizeToFindMatchmakeSession is the max number of PIDs that can be used in FindMatchmakeSessionByParticipant and GetPlayingSession
	MaxPrincipalIDSizeToFindMatchmakeSession int = 300

	// MaxMatchmakeBrowseSizeByParticipant is unknown.
	//
	// Since the value is the same as, and shares a similar name with, MaxPrincipalIDSizeToFindMatchmakeSession
	// this is likely a newer/older version of that constant?
	MaxMatchmakeBrowseSizeByParticipant int = 300

	// MaxMatchmakeSessionByParticipant is the max number of sessions that can be returned by FindMatchmakeSessionByParticipant and GetPlayingSession
	MaxMatchmakeSessionByParticipant int = 1000

	// MaxExtraParticipants is the max value of extraParticipants
	MaxExtraParticipants int = 4

	// MaxP2PSignatureKeyLen is the max size of m_SessionKey
	MaxP2PSignatureKeyLen int = 32

	// ResultRangeAnyOffset when used in MatchmakeSessionSearchCriteria, seems to tell the server to search
	// the entire database for matching sessions and return up to m_uiSize (max 100) sessions in a random
	// order. Otherwise the ResultRange offset is used as normal and sessions are returned in order based on
	// their gathering IDs
	ResultRangeAnyOffset uint64 = 4294967295

	// MatchmakeStringMaxLength is the max length of matchmaking strings
	MatchmakeStringMaxLength int = 256

	// MatchmakeBufferMaxLength is the max size of matchmaking binary data
	MatchmakeBufferMaxLength int = 512
)
