package constants

// RatingLockType represents the type of lock applied to object ratings.
// Locks are applied per-user, not per-object. If a user tries to rate
// an object while a lock is in place, DataStore::OperationNotAllowed is thrown
type RatingLockType uint8

const (

	// RatingLockNone means that the ratings should have no locks
	RatingLockNone RatingLockType = iota

	// RatingLockInterval locks the user from rating the slot again for
	// DataStoreRatingInitParam.periodDuration seconds
	RatingLockInterval

	// RatingLockPeriod locks the user from rating the slot again until
	// a specific date/time. The way this locks expiration is calculated
	// depends on DataStoreRatingInitParam.periodDuration. The value must be
	// a RatingLockPeriod value.
	// If periodDuration is RatingLockPeriodDay1, the expiration day is set
	// to the 1st of the following month regardless of what day that is.
	// If periodDuration is set to the other day-of-the-week values, the
	// expiration day is set to that day of the week of the following week.
	// DataStoreRatingInitParam.periodHour sets the hour, 0-23, of the selected
	// day for when the lock should expire.
	// For example, if periodDuration is set to RatingLockPeriodTue and periodHour
	// is set to 0, and a user rates the slot on March 7th 2024, the lock expiration
	// is set to 12-3-2024 0:00:00 (the following Tuesday at midnight)
	RatingLockPeriod

	// RatingLockPermanent locks the user from rating the slot again forever
	RatingLockPermanent
)
