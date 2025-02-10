package constants

// RatingLockPeriod tells the rating slot locks what day
// the lock should expire
type RatingLockPeriod int16

const (
	// RatingLockPeriodDay1 means the day the lock will expire
	// is the 1st of the following month
	RatingLockPeriodDay1 RatingLockPeriod = -17

	// RatingLockPeriodSun means the day the lock will expire
	// is the Sunday of the following week
	RatingLockPeriodSun RatingLockPeriod = -7

	// RatingLockPeriodSat means the day the lock will expire
	// is the Saturday of the following week
	RatingLockPeriodSat RatingLockPeriod = -6

	// RatingLockPeriodFri means the day the lock will expire
	// is the Friday of the following week
	RatingLockPeriodFri RatingLockPeriod = -5

	// RatingLockPeriodThu means the day the lock will expire
	// is the Thursday of the following week
	RatingLockPeriodThu RatingLockPeriod = -4

	// RatingLockPeriodWed means the day the lock will expire
	// is the Wednesday of the following week
	RatingLockPeriodWed RatingLockPeriod = -3

	// RatingLockPeriodTue means the day the lock will expire
	// is the Tuesday of the following week
	RatingLockPeriodTue RatingLockPeriod = -2

	// RatingLockPeriodMon means the day the lock will expire
	// is the Monday of the following week
	RatingLockPeriodMon RatingLockPeriod = -1
)
