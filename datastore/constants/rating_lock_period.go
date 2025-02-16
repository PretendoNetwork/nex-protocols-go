package constants

// RatingLockPeriod tells the rating slot locks what day
// the lock should expire.
//
// NOTE: Original name is RatingLockPeriod, changed to
// RatingLockPeriodDay as that conflicts with
// RatingLockType.RatingLockPeriod, and these values
// represent the "day" a period lock ends.
type RatingLockPeriodDay int16

const (
	// RatingLockPeriodDay1 means the day the lock will expire
	// is the 1st of the following month
	RatingLockPeriodDay1 RatingLockPeriodDay = -17

	// RatingLockPeriodSun means the day the lock will expire
	// is the Sunday of the following week
	RatingLockPeriodSun RatingLockPeriodDay = -7

	// RatingLockPeriodSat means the day the lock will expire
	// is the Saturday of the following week
	RatingLockPeriodSat RatingLockPeriodDay = -6

	// RatingLockPeriodFri means the day the lock will expire
	// is the Friday of the following week
	RatingLockPeriodFri RatingLockPeriodDay = -5

	// RatingLockPeriodThu means the day the lock will expire
	// is the Thursday of the following week
	RatingLockPeriodThu RatingLockPeriodDay = -4

	// RatingLockPeriodWed means the day the lock will expire
	// is the Wednesday of the following week
	RatingLockPeriodWed RatingLockPeriodDay = -3

	// RatingLockPeriodTue means the day the lock will expire
	// is the Tuesday of the following week
	RatingLockPeriodTue RatingLockPeriodDay = -2

	// RatingLockPeriodMon means the day the lock will expire
	// is the Monday of the following week
	RatingLockPeriodMon RatingLockPeriodDay = -1
)
