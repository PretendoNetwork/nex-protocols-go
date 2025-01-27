package constants

// RankingFilterGroupIndex is used by RankingOrderParam.GroupIndex to select which group to filter by in a score request
type RankingFilterGroupIndex uint8

const (
	// RankingFilterGroupIndex0 indicates RankingOrderParam.GroupNum should be compared to the 0th group.
	RankingFilterGroupIndex0 RankingFilterGroupIndex = iota

	// RankingFilterGroupIndex1 indicates RankingOrderParam.GroupNum should be compared to the 1st group.
	RankingFilterGroupIndex1

	// RankingFilterGroupIndex2 indicates RankingOrderParam.GroupNum should be compared to the 2nd group.
	RankingFilterGroupIndex2

	// RankingFilterGroupIndex3 indicates RankingOrderParam.GroupNum should be compared to the 3rd group.
	RankingFilterGroupIndex3

	// RankingFilterGroupIndexNone indicates that no group filtering should be performed.
	RankingFilterGroupIndexNone = 255
)
