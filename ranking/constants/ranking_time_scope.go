package constants

// RankingTimeScope is used by RankingOrderParam.TimeScope to request that scores only be shown from a certain timeframe.
type RankingTimeScope uint8

const (
	// RankingTimeScopeCustom0 requests only scores from game-specific time scope.
	RankingTimeScopeCustom0 RankingTimeScope = iota

	// RankingTimeScopeCustom1 requests only scores from a second, funnier game-specific time scope.
	RankingTimeScopeCustom1

	// RankingTimeScopeAll requests scores fom all time (no filtering).
	RankingTimeScopeAll
)
