package constants

// RatingLockPeriod tells the rating slot locks what day
// the lock should expire
type RatingLockPeriodType int16

const (
	// RatingLockPeriodDay1 means the day the lock will expire
	// is the 1st of the following month
	RatingLockPeriodDay1 RatingLockPeriodType = -17

	// RatingLockPeriodSun means the day the lock will expire
	// is the Sunday of the following week
	RatingLockPeriodSun RatingLockPeriodType = -7

	// RatingLockPeriodSat means the day the lock will expire
	// is the Saturday of the following week
	RatingLockPeriodSat RatingLockPeriodType = -6

	// RatingLockPeriodFri means the day the lock will expire
	// is the Friday of the following week
	RatingLockPeriodFri RatingLockPeriodType = -5

	// RatingLockPeriodThu means the day the lock will expire
	// is the Thursday of the following week
	RatingLockPeriodThu RatingLockPeriodType = -4

	// RatingLockPeriodWed means the day the lock will expire
	// is the Wednesday of the following week
	RatingLockPeriodWed RatingLockPeriodType = -3

	// RatingLockPeriodTue means the day the lock will expire
	// is the Tuesday of the following week
	RatingLockPeriodTue RatingLockPeriodType = -2

	// RatingLockPeriodMon means the day the lock will expire
	// is the Monday of the following week
	RatingLockPeriodMon RatingLockPeriodType = -1
)
