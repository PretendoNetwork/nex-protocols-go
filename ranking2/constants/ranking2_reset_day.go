package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2ResetDay determines on what day of the week or month Ranking2 season rankings
// will reset.
//
// Note: This is not a real type. This is a bespoke type made for our convenience.
type Ranking2ResetDay uint8

// WriteTo writes the Ranking2ResetDay to the given writable
func (r2rd Ranking2ResetDay) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(r2rd))
}

// ExtractFrom extracts the Ranking2ResetDay value from the given readable
func (r2rd *Ranking2ResetDay) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*r2rd = Ranking2ResetDay(value)
	return nil
}

// String returns a human-readable representation of Ranking2ResetDay.
// This cannot differentiate between types of Ranking2ResetDays
func (r2rd Ranking2ResetDay) String() string {
	return fmt.Sprintf("Ranking2ResetDay(%d)", int(r2rd))
}

// StringWeekday returns a human-readable representation of the Ranking2ResetDay
// when the category reset mode is set to Ranking2ResetModeEveryWeek or Ranking2ResetModeMultiMonthWeekday.
func (r2rd Ranking2ResetDay) StringWeekday() string {
	switch r2rd {
	case Ranking2ResetDayMonday:
		return "Monday"
	case Ranking2ResetDayTuesday:
		return "Tuesday"
	case Ranking2ResetDayWednesday:
		return "Wednesday"
	case Ranking2ResetDayThursday:
		return "Thursday"
	case Ranking2ResetDayFriday:
		return "Friday"
	case Ranking2ResetDaySaturday:
		return "Saturday"
	case Ranking2ResetDaySunday:
		return "Sunday"
	default:
		return fmt.Sprintf("Ranking2ResetDay(%d)", int(r2rd))
	}
}

// StringMonthly returns a human-readable representation of the Ranking2ResetDay
// when the category reset mode is set to Ranking2ResetModeMultiMonth.
func (r2rd Ranking2ResetDay) StringMonthly() string {
	switch r2rd {
	case Ranking2ResetDay1:
		return "1st"
	case Ranking2ResetDay2:
		return "2nd"
	case Ranking2ResetDay3:
		return "3rd"
	case Ranking2ResetDay4:
		return "4th"
	case Ranking2ResetDay5:
		return "5th"
	case Ranking2ResetDay6:
		return "6th"
	case Ranking2ResetDay7:
		return "7th"
	case Ranking2ResetDay8:
		return "8th"
	case Ranking2ResetDay9:
		return "9th"
	case Ranking2ResetDay10:
		return "10th"
	case Ranking2ResetDay11:
		return "11th"
	case Ranking2ResetDay12:
		return "12th"
	case Ranking2ResetDay13:
		return "13th"
	case Ranking2ResetDay14:
		return "14th"
	case Ranking2ResetDay15:
		return "15th"
	case Ranking2ResetDay16:
		return "16th"
	case Ranking2ResetDay17:
		return "17th"
	case Ranking2ResetDay18:
		return "18th"
	case Ranking2ResetDay19:
		return "19th"
	case Ranking2ResetDay20:
		return "20th"
	case Ranking2ResetDay21:
		return "21st"
	case Ranking2ResetDay22:
		return "22nd"
	case Ranking2ResetDay23:
		return "23rd"
	case Ranking2ResetDay24:
		return "24th"
	case Ranking2ResetDay25:
		return "25th"
	case Ranking2ResetDay26:
		return "26th"
	case Ranking2ResetDay27:
		return "27th"
	case Ranking2ResetDay28:
		return "28th"
	default:
		return fmt.Sprintf("Ranking2ResetDay(%d)", int(r2rd))
	}
}

const (
	// Ranking2ResetDayMonday means that season rankings should reset every Monday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Monday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDayMonday Ranking2ResetDay = 0

	// Ranking2ResetDayTuesday means that season rankings should reset every Tuesday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Tuesday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDayTuesday Ranking2ResetDay = 1

	// Ranking2ResetDayWednesday means that season rankings should reset every Wednesday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Wednesday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDayWednesday Ranking2ResetDay = 2

	// Ranking2ResetDayThursday means that season rankings should reset every Thursday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Thursday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDayThursday Ranking2ResetDay = 3

	// Ranking2ResetDayFriday means that season rankings should reset every Friday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Friday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDayFriday Ranking2ResetDay = 4

	// Ranking2ResetDaySaturday means that season rankings should reset every Saturday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Saturday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDaySaturday Ranking2ResetDay = 5

	// Ranking2ResetDaySunday means that season rankings should reset every Sunday
	// when the category reset mode is set to Ranking2ResetModeEveryWeek. This also means
	// that season rankings should reset on the first Sunday of the month when the category
	// reset mode is set to Ranking2ResetModeMultiMonthWeekday.
	Ranking2ResetDaySunday Ranking2ResetDay = 6
)

const (
	// Ranking2ResetDay1 means that season rankings should reset on the 1st of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay1 Ranking2ResetDay = 1

	// Ranking2ResetDay2 means that season rankings should reset on the 2nd of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay2 Ranking2ResetDay = 2

	// Ranking2ResetDay3 means that season rankings should reset on the 3rd of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay3 Ranking2ResetDay = 3

	// Ranking2ResetDay4 means that season rankings should reset on the 4th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay4 Ranking2ResetDay = 4

	// Ranking2ResetDay5 means that season rankings should reset on the 5th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay5 Ranking2ResetDay = 5

	// Ranking2ResetDay6 means that season rankings should reset on the 6th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay6 Ranking2ResetDay = 6

	// Ranking2ResetDay7 means that season rankings should reset on the 7th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay7 Ranking2ResetDay = 7

	// Ranking2ResetDay8 means that season rankings should reset on the 8th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay8 Ranking2ResetDay = 8

	// Ranking2ResetDay9 means that season rankings should reset on the 9th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay9 Ranking2ResetDay = 9

	// Ranking2ResetDay10 means that season rankings should reset on the 10th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay10 Ranking2ResetDay = 10

	// Ranking2ResetDay11 means that season rankings should reset on the 11th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay11 Ranking2ResetDay = 11

	// Ranking2ResetDay12 means that season rankings should reset on the 12th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay12 Ranking2ResetDay = 12

	// Ranking2ResetDay13 means that season rankings should reset on the 13th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay13 Ranking2ResetDay = 13

	// Ranking2ResetDay14 means that season rankings should reset on the 14th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay14 Ranking2ResetDay = 14

	// Ranking2ResetDay15 means that season rankings should reset on the 15th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay15 Ranking2ResetDay = 15

	// Ranking2ResetDay16 means that season rankings should reset on the 16th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay16 Ranking2ResetDay = 16

	// Ranking2ResetDay17 means that season rankings should reset on the 17th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay17 Ranking2ResetDay = 17

	// Ranking2ResetDay18 means that season rankings should reset on the 18th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay18 Ranking2ResetDay = 18

	// Ranking2ResetDay19 means that season rankings should reset on the 19th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay19 Ranking2ResetDay = 19

	// Ranking2ResetDay20 means that season rankings should reset on the 20th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay20 Ranking2ResetDay = 20

	// Ranking2ResetDay21 means that season rankings should reset on the 21st of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay21 Ranking2ResetDay = 21

	// Ranking2ResetDay22 means that season rankings should reset on the 22nd of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay22 Ranking2ResetDay = 22

	// Ranking2ResetDay23 means that season rankings should reset on the 23rd of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay23 Ranking2ResetDay = 23

	// Ranking2ResetDay24 means that season rankings should reset on the 24th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay24 Ranking2ResetDay = 24

	// Ranking2ResetDay25 means that season rankings should reset on the 25th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay25 Ranking2ResetDay = 25

	// Ranking2ResetDay26 means that season rankings should reset on the 26th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay26 Ranking2ResetDay = 26

	// Ranking2ResetDay27 means that season rankings should reset on the 27th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay27 Ranking2ResetDay = 27

	// Ranking2ResetDay28 means that season rankings should reset on the 28th of the
	// the month in an enabled month when the category reset mode is set to Ranking2ResetModeMultiMonth
	Ranking2ResetDay28 Ranking2ResetDay = 28
)
