package constants

// * Miscellaneous constants

const (
	// MaxPeriod is the maximum period an object expiration can have
	MaxPeriod uint16 = 365

	// MaxMetaBinSize is the maximum size of an object MetaBinary
	MaxMetaBinSize uint32 = 1024

	// DatastorePermissionRecipientIDsMax is the maximum number of
	// recipient IDs allowed for permissions
	DatastorePermissionRecipientIDsMax uint32 = 100

	// InvalidDataID represents an invalid DataID. This usually
	// indicates that the field should be ignored
	InvalidDataID uint64 = 0

	// InvalidDataType represents an invalid DataType. This usually
	// indicates that the field should be ignored
	InvalidDataType uint16 = 65535

	// InvalidPassword represents an invalid password. This usually
	// indicates that the field should be ignored or that a password
	// was not set
	InvalidPassword uint64 = 0

	// MaxNameLength is the maximum object name length
	MaxNameLength uint32 = 64

	// MaxSearchResultSize is the maximum number of object search results
	MaxSearchResultSize uint32 = 100

	// MaxSearchAnyResultSize is the maximum number of search-any results
	MaxSearchAnyResultSize uint32 = 20

	// MaxSearchDataTypeSize is the maximum number of DataTypes
	// allowed in DataStoreSearchParam.dataTypes
	MaxSearchDataTypeSize uint32 = 10

	// NumTagSlot is the maximum number of tags an object can have.
	// Slots may use IDs 0-15
	NumTagSlot uint32 = 16

	// RatingSlotMax is the maximum rating slot index
	RatingSlotMax uint32 = 15

	// NumRatingSlot is the maximum number of rating slots an
	// object can have. Slots may use IDs 0-15
	NumRatingSlot uint32 = 16

	// MaxTagLength is the maximum length of an object tag
	MaxTagLength uint32 = 24

	// DefaultPeriod is the default object expiration period
	DefaultPeriod uint16 = 90

	// DefaultHTTPThreadPriority has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultHTTPThreadPriority uint32 = 16

	// DefaultRelayBufferSize has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultRelayBufferSize uint32 = 16384

	// DefaultHTTPBufferSize has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultHTTPBufferSize uint32 = 32768

	// DefaultDataTransferTimeoutBytesPerSecond has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultDataTransferTimeoutBytesPerSecond uint32 = 167

	// DefaultDataTransferMinimumTimeout has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultDataTransferMinimumTimeout int = 60000

	// DefaultHTTPSendSocketBufferSize has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultHTTPSendSocketBufferSize uint32 = 65536

	// DefaultHTTPRecvSocketBufferSize has a currently unknown use.
	// Likely only used by clients when downloading/uploading
	// objects to the storage server?
	DefaultHTTPRecvSocketBufferSize uint32 = 65536

	// InvalidPersistenceSlotID represents an invalid persistence slot.
	// This indicates that an object should not be persisted
	InvalidPersistenceSlotID uint16 = 65535

	// NumPersistenceSlot is the maximum number of objects each user
	// may persist. Slots may use IDs 0-15
	NumPersistenceSlot uint16 = 16

	// BatchProcessingCapacityPostObject is the maximum number of
	// objects which the server create at one time. For example,
	// RateObjectsWithPosting
	BatchProcessingCapacityPostObject uint32 = 16

	// BatchProcessingCapacity is the maximum number of objects which
	// the server can process in a single request. For example, GetMetas
	BatchProcessingCapacity uint32 = 100

	// ResultRangeAnyOffset has a currently unknown use. Likely means
	// that the server should pick any random offset?
	ResultRangeAnyOffset uint64 = 4294967295

	// ResultRangeDefaultSize is the default size of ResultRange
	ResultRangeDefaultSize uint32 = 20
)
