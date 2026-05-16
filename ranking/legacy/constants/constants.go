package constants

const (
	// MaxCommonDataSize is the maximum length a users common data may be.
	//
	// Note: This is a guess based on the NEX 3 Ranking constants
	MaxCommonDataSize int = 255

	// CommonDataSizeNEX1 is the length a users common data must be on the NEX 1 implementation
	CommonDataSizeNEX1 int = 20

	// MaxRangeRankingOrder is the maximum value an offset may be when
	// requesting ranking data.
	//
	// Note: This is a guess based on the NEX 3 Ranking constants
	MaxRangeRankingOrder uint32 = 1000

	// MaxAccurateOrder has an unknown use. This likely is used to indicate
	// the largest number of ranking results that have accurate ranking orders?
	//
	// For example, if 4999 rankings are returned, the client knows that this
	// data has accurate rankings. but if 5001 are returned, they may not be?
	//
	// Note: This is a guess based on the NEX 3 Ranking constants
	MaxAccurateOrder uint32 = 5000

	// NumRankingDataScores is the number of scores that a ranking data must have.
	//
	// Note: The name is a guess from "NumMatchmakeSessionAttributes"
	NumRankingDataScores int = 2

	// NumRankingDataScores is the number of categories that a ranking data can have on NEX 1.
	//
	// Note: The name is a guess from "NumMatchmakeSessionAttributes"
	NumRankingDataCategories int = 1
)
