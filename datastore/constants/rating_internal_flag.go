package constants

// RatingInternalFlag indicates whether or not a minimum or
// maximum rating value should be set
type RatingInternalFlag uint8

const (

	// RatingInternalFlagUseRangeMin means that the DataStoreRatingInitParam.rangeMin
	// value should be respected. If this flag is not set, the minimum value is ignored.
	// If this flag is set and a rating is below the minimum, `DataStore::InvalidArgument`
	// is thrown
	RatingInternalFlagUseRangeMin RatingInternalFlag = 0x2

	// RatingInternalFlagUseRangeMax means that the DataStoreRatingInitParam.rangeMax
	// value should be respected. If this flag is not set, the maximum value is ignored.
	// If this flag is set and a rating is above the maximum, `DataStore::InvalidArgument`
	// is thrown
	RatingInternalFlagUseRangeMax RatingInternalFlag = 0x4
)
