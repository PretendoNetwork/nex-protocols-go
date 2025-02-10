package constants

// RatingFlag indicates how the server should handle
// user ratings for object rating slots
type RatingFlag uint8

const (
	// RatingFlagModifiable means that if a user rates an object
	// rating slot more than once, update the existing rating
	// rather than create a new one. If this flag is set, ratings
	// to slots with existing ratings from the user do not count
	// towards the rating count. If this flag is not set, ratings
	// to slots with existing ratings from the user are treated
	// as separate ratings and the rating count is incremented
	RatingFlagModifiable RatingFlag = 0x4

	// RatingFlagRoundMinus means that any rating value smaller
	// than 0 is rounded up to 0 before being handled
	RatingFlagRoundMinus RatingFlag = 0x8

	// RatingFlagDisableSelfRating means that a user cannot rate
	// an object that they own. If this flag is set and a user
	// tries to rate an object they own, DataStore::OperationNotAllowed is thrown
	RatingFlagDisableSelfRating RatingFlag = 0x10
)
