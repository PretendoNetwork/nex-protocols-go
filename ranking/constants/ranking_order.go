package constants

// RankingOrderBy is used in RankingScoreData.OrderBy to set the "golf scoring" mode for a category.
type RankingOrderBy uint8

const (
	RankingOrderByAscending RankingOrderBy = iota
	RankingOrderByDescending
)
