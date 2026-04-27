package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SearchSortColumn tells the server which database column to use as the input
// for the ordering of returned object searches
type SearchSortColumn uint8

// WriteTo writes the SearchSortColumn to the given writable
func (ssc SearchSortColumn) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(ssc))
}

// ExtractFrom extracts the SearchSortColumn value from the given readable
func (ssc *SearchSortColumn) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*ssc = SearchSortColumn(value)
	return nil
}

// String returns a human-readable representation of the SearchSortColumn.
func (ssc SearchSortColumn) String() string {
	switch ssc {
	case SearchSortColumnDataID:
		return "DataID"
	case SearchSortColumnSize:
		return "Size"
	case SearchSortColumnNameAlphabetical:
		return "NameAlphabetical"
	case SearchSortColumnDataType:
		return "DataType"
	case SearchSortColumnReferredCount:
		return "ReferredCount"
	case SearchSortColumnCreatedTime:
		return "CreatedTime"
	case SearchSortColumnUpdatedTime:
		return "UpdatedTime"
	case SearchSortColumnRating0:
		return "Rating0"
	case SearchSortColumnRating1:
		return "Rating1"
	case SearchSortColumnRating2:
		return "Rating2"
	case SearchSortColumnRating3:
		return "Rating3"
	case SearchSortColumnRating4:
		return "Rating4"
	case SearchSortColumnRating5:
		return "Rating5"
	case SearchSortColumnRating6:
		return "Rating6"
	case SearchSortColumnRating7:
		return "Rating7"
	case SearchSortColumnRating8:
		return "Rating8"
	case SearchSortColumnRating9:
		return "Rating9"
	case SearchSortColumnRating10:
		return "Rating10"
	case SearchSortColumnRating11:
		return "Rating11"
	case SearchSortColumnRating12:
		return "Rating12"
	case SearchSortColumnRating13:
		return "Rating13"
	case SearchSortColumnRating14:
		return "Rating14"
	case SearchSortColumnRating15:
		return "Rating15"
	case SearchSortColumnRatingAverage0:
		return "RatingAverage0"
	case SearchSortColumnRatingAverage1:
		return "RatingAverage1"
	case SearchSortColumnRatingAverage2:
		return "RatingAverage2"
	case SearchSortColumnRatingAverage3:
		return "RatingAverage3"
	case SearchSortColumnRatingAverage4:
		return "RatingAverage4"
	case SearchSortColumnRatingAverage5:
		return "RatingAverage5"
	case SearchSortColumnRatingAverage6:
		return "RatingAverage6"
	case SearchSortColumnRatingAverage7:
		return "RatingAverage7"
	case SearchSortColumnRatingAverage8:
		return "RatingAverage8"
	case SearchSortColumnRatingAverage9:
		return "RatingAverage9"
	case SearchSortColumnRatingAverage10:
		return "RatingAverage10"
	case SearchSortColumnRatingAverage11:
		return "RatingAverage11"
	case SearchSortColumnRatingAverage12:
		return "RatingAverage12"
	case SearchSortColumnRatingAverage13:
		return "RatingAverage13"
	case SearchSortColumnRatingAverage14:
		return "RatingAverage14"
	case SearchSortColumnRatingAverage15:
		return "RatingAverage15"
	default:
		return fmt.Sprintf("SearchSortColumn(%d)", int(ssc))
	}
}

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
