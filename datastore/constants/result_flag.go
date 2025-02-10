package constants

// ResultFlag tells the server what fields to populate in responses
// to object searches
type ResultFlag uint8

const (
	// ResultFlagTags means the object tags should be populated
	ResultFlagTags ResultFlag = 0x1

	// ResultFlagRatings means the object ratings should be populated
	ResultFlagRatings ResultFlag = 0x2

	// ResultFlagMetabinary means the object MetaBinary should be populated
	ResultFlagMetabinary ResultFlag = 0x4

	// ResultFlagPermittedIDs means the object permissions should be populated
	ResultFlagPermittedIDs ResultFlag = 0x8
)
