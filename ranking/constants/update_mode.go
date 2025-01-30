package constants

// UpdateMode is used by RankingScoreData.UpdateMode to control if worse scores should be discarded or not.
type UpdateMode uint8

const (
	// UpdateModeNormal will only commit the updated score if it's better than the old score in that category.
	// "better" is determined by the ranking order of the category.
	UpdateModeNormal UpdateMode = iota

	// UpdateModeDeleteOld will always overwrite any old score, even if the new one is worse.
	UpdateModeDeleteOld
)
