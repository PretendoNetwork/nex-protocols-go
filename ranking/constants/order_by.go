package constants

// OrderBy is used in RankingScoreData.OrderBy to set the "golf scoring" mode for a category.
type OrderBy uint8

const (
	OrderByAscending OrderBy = iota
	OrderByDescending
)
