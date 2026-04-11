package constants

const (
	// MaxCommonDataSize is the maximum length a users common data may be.
	MaxCommonDataSize int = 255

	// MaxRangeRankingOrder is the maximum value an offset may be when
	// requesting ranking data.
	MaxRangeRankingOrder uint32 = 1000

	// MaxAccurateOrder has an unknown use. This likely is used to indicate
	// the largest number of ranking results that have accurate ranking orders?
	//
	// For example, if 4999 rankings are returned, the client knows that this
	// data has accurate rankings. but if 5001 are returned, they may not be?
	MaxAccurateOrder uint32 = 5000
)
