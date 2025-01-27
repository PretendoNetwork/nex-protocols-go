package constants

// RankingMode represents the selection of who will be included on the leaderboard (global, friends, nearby etc.).
type RankingMode uint32

const (
	// RankingModeRange retrieves the entire global leaderboard up to 1000 entries.
	// `RankingOrderParam.offset` is the base rank to offset from (0 being world record), no higher than 1000 (MAX_RANGE_RANKING_ORDER?).
	// PIDs and unique IDs are ignored.
	RankingModeRange RankingMode = iota

	// RankingModeNear retrieves the selected users rankings and those surrounding the selected user,
	// placing the selected user in the middle. For example if `RankingOrderParam.length` is 11, then
	// the selected user's ranking would be in the 5th index (6th place), with 5 rankings on both sides.
	// The selected user may not be in the middle if in first or last place.
	// "Selected user" can mean both the connected user OR specific user(s). If a user is specified
	// (PID/Unique ID), that becomes "selected user". Otherwise this is the connected user.
	// More than 1 "selected user" may be selected in methods such as `GetRankingByPIDList`
	// `RankingOrderParam.offset` is ignored.
	RankingModeNear

	// RankingModeFriendRange functions identically to RankingModeRange, but only returns rankings
	// of friends and the connected user. All data from all unique IDs for the connected user and
	// friends are retrieved.
	// If the connected user has no friends, functions identically to RankingModeUser.
	// PIDs and unique IDs are ignored.
	RankingModeFriendRange

	// RankingModeFriendNear retrieves the connected users rankings and the rankings of friends surrounding
	// the connected user, placing the connected user in the middle. For example if `RankingOrderParam.length`
	// is 11, then the connected user's ranking would be in the 5th index (6th place), with 5 friends on both sides.
	// The connected user may not be in the middle if in first or last place.
	// If the connected user has more than 1 unique ID, only respect the ones requested (if any).
	// `RankingOrderParam.offset` is ignored.
	// PIDs and `RankingOrderParam.offset` are ignored.
	RankingModeFriendNear

	// RankingModeUser retrieves all the ranking data for the connected user, including all unique IDs.
	// `RankingOrderParam.offset` is ignored.
	RankingModeUser
)
