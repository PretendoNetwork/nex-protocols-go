package constants

// SearchSortColumn tells the server which database column to use as the input
// for the ordering of returned object searches
type SearchSortColumn uint8

const (
	// SearchSortColumnDataID means the objects should be sorted based on the objects DataIDs
	SearchSortColumnDataID SearchSortColumn = iota

	// SearchSortColumnSize means the objects should be sorted based on the objects sizes
	SearchSortColumnSize

	// SearchSortColumnNameAlphabetical means the objects should be sorted based on the objects names
	SearchSortColumnNameAlphabetical

	// SearchSortColumnDataType means the objects should be sorted based on the objects data types
	SearchSortColumnDataType

	// SearchSortColumnReferredCount means the objects should be sorted based on the objects referred counts
	SearchSortColumnReferredCount // * Mostly a guess. Modern DataStore seems to not use refer counts anymore? Guessed based on similar positions in other enums

	// SearchSortColumnCreatedTime means the objects should be sorted based on the objects created times
	SearchSortColumnCreatedTime

	// SearchSortColumnUpdatedTime means the objects should be sorted based on the objects update times
	SearchSortColumnUpdatedTime
)

const (
	// SearchSortColumnRating0 means the objects should be sorted based on total combined ratings of slot 0
	SearchSortColumnRating0 SearchSortColumn = iota + 64

	// SearchSortColumnRating1 means the objects should be sorted based on total combined ratings of slot 1
	SearchSortColumnRating1

	// SearchSortColumnRating2 means the objects should be sorted based on total combined ratings of slot 2
	SearchSortColumnRating2

	// SearchSortColumnRating3 means the objects should be sorted based on total combined ratings of slot 3
	SearchSortColumnRating3

	// SearchSortColumnRating4 means the objects should be sorted based on total combined ratings of slot 4
	SearchSortColumnRating4

	// SearchSortColumnRating5 means the objects should be sorted based on total combined ratings of slot 5
	SearchSortColumnRating5

	// SearchSortColumnRating6 means the objects should be sorted based on total combined ratings of slot 6
	SearchSortColumnRating6

	// SearchSortColumnRating7 means the objects should be sorted based on total combined ratings of slot 7
	SearchSortColumnRating7

	// SearchSortColumnRating8 means the objects should be sorted based on total combined ratings of slot 8
	SearchSortColumnRating8

	// SearchSortColumnRating9 means the objects should be sorted based on total combined ratings of slot 9
	SearchSortColumnRating9

	// SearchSortColumnRating10 means the objects should be sorted based on total combined ratings of slot 10
	SearchSortColumnRating10

	// SearchSortColumnRating11 means the objects should be sorted based on total combined ratings of slot 11
	SearchSortColumnRating11

	// SearchSortColumnRating12 means the objects should be sorted based on total combined ratings of slot 12
	SearchSortColumnRating12

	// SearchSortColumnRating13 means the objects should be sorted based on total combined ratings of slot 13
	SearchSortColumnRating13

	// SearchSortColumnRating14 means the objects should be sorted based on total combined ratings of slot 14
	SearchSortColumnRating14

	// SearchSortColumnRating15 means the objects should be sorted based on total combined ratings of slot 15
	SearchSortColumnRating15
)

const (
	// SearchSortColumnRatingAverage0 means the objects should be sorted based on the average value of ratings of slot 0
	SearchSortColumnRatingAverage0 SearchSortColumn = iota + 96

	// SearchSortColumnRatingAverage1 means the objects should be sorted based on the average value of ratings of slot 1
	SearchSortColumnRatingAverage1

	// SearchSortColumnRatingAverage2 means the objects should be sorted based on the average value of ratings of slot 2
	SearchSortColumnRatingAverage2

	// SearchSortColumnRatingAverage3 means the objects should be sorted based on the average value of ratings of slot 3
	SearchSortColumnRatingAverage3

	// SearchSortColumnRatingAverage4 means the objects should be sorted based on the average value of ratings of slot 4
	SearchSortColumnRatingAverage4

	// SearchSortColumnRatingAverage5 means the objects should be sorted based on the average value of ratings of slot 5
	SearchSortColumnRatingAverage5

	// SearchSortColumnRatingAverage6 means the objects should be sorted based on the average value of ratings of slot 6
	SearchSortColumnRatingAverage6

	// SearchSortColumnRatingAverage7 means the objects should be sorted based on the average value of ratings of slot 7
	SearchSortColumnRatingAverage7

	// SearchSortColumnRatingAverage8 means the objects should be sorted based on the average value of ratings of slot 8
	SearchSortColumnRatingAverage8

	// SearchSortColumnRatingAverage9 means the objects should be sorted based on the average value of ratings of slot 9
	SearchSortColumnRatingAverage9

	// SearchSortColumnRatingAverage10 means the objects should be sorted based on the average value of ratings of slot 10
	SearchSortColumnRatingAverage10

	// SearchSortColumnRatingAverage11 means the objects should be sorted based on the average value of ratings of slot 11
	SearchSortColumnRatingAverage11

	// SearchSortColumnRatingAverage12 means the objects should be sorted based on the average value of ratings of slot 12
	SearchSortColumnRatingAverage12

	// SearchSortColumnRatingAverage13 means the objects should be sorted based on the average value of ratings of slot 13
	SearchSortColumnRatingAverage13

	// SearchSortColumnRatingAverage14 means the objects should be sorted based on the average value of ratings of slot 14
	SearchSortColumnRatingAverage14

	// SearchSortColumnRatingAverage15 means the objects should be sorted based on the average value of ratings of slot 15
	SearchSortColumnRatingAverage15
)
