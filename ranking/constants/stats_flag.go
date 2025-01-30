package constants

// StatsFlag is a bitmask used by GetRanking to request the inclusion of different aggregate stats
type StatsFlag uint8

const (
	// StatsFlagTotal requests the total of the stats
	StatsFlagTotal = 0x1

	// StatsFlagSum requests the sum of the stats
	StatsFlagSum = 0x2

	// StatsFlagMin requests the minimum stat
	StatsFlagMin = 0x4

	// StatsFlagMax requests the maximum stat
	StatsFlagMax = 0x8

	// StatsFlagAverage requests the average of the stats
	StatsFlagAverage = 0x10
)
