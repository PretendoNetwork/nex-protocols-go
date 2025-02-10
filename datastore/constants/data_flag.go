package constants

// DataFlag sets different configuration flags for uploaded objects.
// Stored in the objects DataStoreMetaInfo.flag field
type DataFlag uint8

const (
	// DataFlagNone means no extra configurations
	DataFlagNone DataFlag = 0x0

	// DataFlagNeedReview means that the object status should be set to
	// DataStatusPending and can only be seen by the owner until reviewed.
	// Trying to query for objects under review will result in DataStore::UnderReviewing
	DataFlagNeedReview DataFlag = 0x1

	// DataFlagPeriodFromLastReferred means that the object should have it's expiration
	// time increased by the objects DataStoreMetaInfo.period days when directly interacted
	// with (searched for by ID, touched with TouchObject, etc.).
	// If this flag is not set, objects will expire after DataStoreMetaInfo.period days and
	// be removed from the game/storage server
	DataFlagPeriodFromLastReferred DataFlag = 0x2

	// DataFlagUseReadLock has an unknown use. Seems to do nothing, likely unused?
	DataFlagUseReadLock DataFlag = 0x4

	// DataFlagUseNotificationOnPost means that when an object is created and it's
	// access permission is not set to PermissionPublic or PermissionPrivate, then
	// all users who have access permission for the object receive a DataStore
	// notification
	DataFlagUseNotificationOnPost DataFlag = 0x8

	// DataFlagUseNotificationOnPost means that when an existing object is updated
	// and it's access permission is not set to PermissionPublic or PermissionPrivate,
	// then all users who have access permission for the object receive a DataStore
	// notification
	DataFlagUseNotificationOnUpdate DataFlag = 0x10

	// DataFlagNotUseFileServer means the entry in the game server has no physical object in the
	// storage (s3) server. Setting this flag when uploading an object will throw DataStore::InvalidArgument.
	// Flag is automatically set on entries when uploading just a MetaBinary
	DataFlagNotUseFileServer DataFlag = 0x20

	// DataFlagNeedCompletion means "completion" must be called when uploading/updating an object.
	// If uploading an object, CompletePostObjectV1/CompletePostObject MUST be called.
	// If updating an object, CompleteUpdateObject MUST be called. Until an object is completed
	// it is treated as non-existent (DataStore::NotFound)
	DataFlagNeedCompletion DataFlag = 0x40
)
