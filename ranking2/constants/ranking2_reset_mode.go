package constants

// Ranking2ResetMode determines when and how to reset ranking data. Ranking resets
// begin a new "season" in Ranking2.
type Ranking2ResetMode uint8

const (
	// Ranking2ResetModeNothing means that ranking data never resets.
	Ranking2ResetModeNothing Ranking2ResetMode = 0

	// Ranking2ResetModeEveryDay means that ranking data resets every day.
	// The `resetHour` value determines when in the day to reset.
	Ranking2ResetModeEveryDay Ranking2ResetMode = 1

	// Ranking2ResetModeEveryWeek means that ranking data resets every week
	// on a specified day. `resetDay` is used to determine the day of the week,
	// and `resetHour` determines when in the day to reset.
	//
	// Note: The exact value range for `resetDay` is not known in this mode.
	// It is assumed to be treated as an enum, using standard 0-based indexes.
	// That would make the values 0-6 assuming typical Monday-Sunday ordering.
	Ranking2ResetModeEveryWeek Ranking2ResetMode = 2

	// Ranking2ResetModeMultiMonth means that ranking data resets every enabled
	// month. `resetMonth` is used to determine which months are enabled for resets.
	// This value appears to be a 12-bit set of flags, where each bit represents
	// one of the 12 months. If a bit is set, the rankings reset that month. The
	// `resetDay` is used to determine the day of the week, and `resetHour`
	// determines when in the day to reset.
	//
	// Note: The order of the `resetMonth` bits is not currently known. It can
	// safely be assumed to go in month order, however. Starting with January
	// and ending with December where January is the LSB. The value range for
	// `resetDay` is also not known, however as it represents a calendar date it
	// can be safely assumed to begin at 1. since not all months have the same
	// number of days, it can be safely assumed that the upper limit would be the
	// most amount of days all months have, making the range 1-28. This does, however,
	// mean that months with more than 28 days cannot have rankings reset on those days.
	Ranking2ResetModeMultiMonth Ranking2ResetMode = 3

	// Ranking2ResetModeMultiMonthWeekday means that ranking data resets every enabled
	// month on a specific day of the week? This seems to be a mix of `Ranking2ResetModeEveryWeek`
	// and `Ranking2ResetModeMultiMonth`. `resetMonth` is likely still used to determine
	// what months are enabled for resets, but `resetDay` now refers to a day of the week
	// rather than a calendar date. Given that this modes name says "Weekday", and not
	// "EveryWeek", it's likely safe to assume it is NOT every week? This may indicate
	// that the reset happens on the first weekday set in `resetDay` of the enabled months.
	// For example, resetting rankings on the first Monday every other month?
	//
	// Note: This functionality is entirely as guess, based on it's name and the
	// functionality of the other modes. It could be entirely incorrect.
	Ranking2ResetModeMultiMonthWeekday Ranking2ResetMode = 4
)
