package constants

// RankingOrderCalculation is used in RankingOrderParam.OrderCalculation to control how ties are handled.
type RankingOrderCalculation uint8

const (
	// RankingOrderCalculation113 requests standard "1224" competition ranking.
	RankingOrderCalculation113 RankingOrderCalculation = iota

	// RankingOrderCalculation123 requests strictly ordinal "1234" ranking, with ties broken first by update time
	// (earlier is better) then user PID (lower is better).
	RankingOrderCalculation123
)
