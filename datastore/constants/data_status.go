package constants

// DataStatus indicates the availablity status of an object.
// If an object has a status other than DataStatusNone, it is
// not available to anyone besides the owner and only under
// certain circumstances
type DataStatus uint8

const (
	// DataStatusNone means the object is uploaded and
	// visible to those who have access permissions
	DataStatusNone DataStatus = 0

	// DataStatusPending means the object is uploaded but
	// pending review. The can only be seen by the owner until reviewed.
	// Trying to query for objects under review will result in DataStore::UnderReviewing
	DataStatusPending DataStatus = 2

	// DataStatusRejected means the object is uploaded but
	// is not available to the public. Likely means rejected
	// from review
	DataStatusRejected DataStatus = 5
)
