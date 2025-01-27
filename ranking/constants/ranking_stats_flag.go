package constants

// RankingStatsFlag is a bitmask used by GetRanking to request the inclusion of different aggregate stats
type RankingStatsFlag uint8

const (
	// RankingStatsFlagTotal requests the total of the stats
	RankingStatsFlagTotal = 0x1

	// RankingStatsFlagSum requests the sum of the stats
	RankingStatsFlagSum = 0x2

	// RankingStatsFlagMin requests the minimum stat
	RankingStatsFlagMin = 0x4

	// RankingStatsFlagMax requests the maximum stat
	RankingStatsFlagMax = 0x8

	// RankingStatsFlagAverage requests the average of the stats
	RankingStatsFlagAverage = 0x10
)
