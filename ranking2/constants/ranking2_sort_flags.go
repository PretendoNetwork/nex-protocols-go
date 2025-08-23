package constants

// Ranking2SortFlags determines how ranking results should be ordered.
type Ranking2SortFlags uint32

const (
	// Ranking2SortFlagsNothing means results should be returned in no
	// specific order.
	Ranking2SortFlagsNothing Ranking2SortFlags = 0

	// Ranking2SortFlagsAscending means results should be returned in no
	// ascending order?
	//
	// Note: This is a guess based on some light test behavior and the fact
	// that there is no other ordering flag like this for the request it's
	// used in. The real name of this is unknown. And functionality not 100%
	// confirmed.
	Ranking2SortFlagsAscending Ranking2SortFlags = 1

	// Ranking2SortFlagsDescending means results should be returned in no
	// descending order?
	//
	// Note: This is a guess based on some light test behavior and the fact
	// that there is no other ordering flag like this for the request it's
	// used in. The real name of this is unknown. And functionality not 100%
	// confirmed.
	Ranking2SortFlagsDescending Ranking2SortFlags = 2

	// Ranking2SortFlagsMoveToTopInTie means in the event of a tie between
	// another user and the caller, the caller is ranked better than the
	// other user.
	Ranking2SortFlagsMoveToTopInTie Ranking2SortFlags = 4
)
