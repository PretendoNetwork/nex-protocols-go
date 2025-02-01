package constants

// ModificationFlag is used by RankingChangeAttributesParam.ModificationFlag to set which fields should be
// updated.
type ModificationFlag uint8

const (
	// ModificationFlagNone indicates that no updates should occur.
	ModificationFlagNone ModificationFlag = 0x0

	// ModificationFlagGroup0 indicates that the group at the 0th group should be updated.
	ModificationFlagGroup0 = 0x1

	// ModificationFlagGroup1 indicates that the group at the 1st group should be updated.
	ModificationFlagGroup1 = 0x2

	// ModificationFlagGroup2 indicates that the group at the 2nd group should be updated.
	// Unused, only 2 groups are supported
	ModificationFlagGroup2 = 0x4

	// ModificationFlagGroup3 indicates that the group at the 3rd group should be updated.
	// Unused, only 2 groups are supported
	ModificationFlagGroup3 = 0x8

	// ModificationFlagParam indicates that the param should be updated.
	ModificationFlagParam = 0x10
)
