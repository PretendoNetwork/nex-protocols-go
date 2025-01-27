package constants

// RankingUpdateMode is used by RankingScoreData.UpdateMode to control if worse scores should be discarded or not.
type RankingUpdateMode uint8

const (
	// RankingUpdateModeNormal will only commit the updated score if it's better than the old score in that category.
	// "better" is determined by the ranking order of the category.
	RankingUpdateModeNormal RankingUpdateMode = iota

	// RankingUpdateModeDeleteOld will always overwrite any old score, even if the new one is worse.
	RankingUpdateModeDeleteOld
)
