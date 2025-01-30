package constants

// TimeScope is used by RankingOrderParam.TimeScope to request that scores only be shown from a certain timeframe.
type TimeScope uint8

const (
	// TimeScopeCustom0 requests only scores from game-specific time scope.
	TimeScopeCustom0 TimeScope = iota

	// TimeScopeCustom1 requests only scores from a second, funnier game-specific time scope.
	TimeScopeCustom1

	// TimeScopeAll requests scores fom all time (no filtering).
	TimeScopeAll
)
