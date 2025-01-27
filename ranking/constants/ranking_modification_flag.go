package constants

// RankingModificationFlag is used by RankingChangeAttributesParam.ModificationFlag to set which fields should be
// updated.
type RankingModificationFlag uint8

const (
	// RankingModificationFlagNone indicates that no updates should occur.
	RankingModificationFlagNone RankingModificationFlag = 0x0

	// RankingModificationFlagGroup0 indicates that the group at the 0th group should be updated.
	RankingModificationFlagGroup0 = 0x1

	// RankingModificationFlagGroup1 indicates that the group at the 1st group should be updated.
	RankingModificationFlagGroup1 = 0x2

	// RankingModificationFlagGroup2 indicates that the group at the 2nd group should be updated.
	// Unused, only 2 groups are supported
	RankingModificationFlagGroup2 = 0x4

	// RankingModificationFlagGroup3 indicates that the group at the 3rd group should be updated.
	// Unused, only 2 groups are supported
	RankingModificationFlagGroup3 = 0x8

	// RankingModificationFlagParam indicates that the param should be updated.
	RankingModificationFlagParam = 0x10
)
