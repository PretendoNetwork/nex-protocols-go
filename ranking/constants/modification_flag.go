package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ModificationFlag is used by RankingChangeAttributesParam.ModificationFlag to set which fields should be
// updated.
type ModificationFlag uint8

// WriteTo writes the ModificationFlag to the given writable
func (mf ModificationFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(mf))
}

// ExtractFrom extracts the ModificationFlag value from the given readable
func (mf *ModificationFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*mf = ModificationFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (mf ModificationFlag) HasFlag(flag ModificationFlag) bool {
	return mf&flag == flag
}

// HasFlag checks if all given flags are set
func (mf ModificationFlag) HasFlags(flags ...ModificationFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if mf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the ModificationFlag bitmask.
// Multiple flags are joined with "|", e.g. "Group0|Group1|Param".
// Returns "None" if no flags are set.
func (mf ModificationFlag) String() string {
	if mf == ModificationFlagNone {
		return "None"
	}

	flags := []struct {
		flag ModificationFlag
		name string
	}{
		{ModificationFlagGroup0, "Group0"},
		{ModificationFlagGroup1, "Group1"},
		{ModificationFlagGroup2, "Group2"},
		{ModificationFlagGroup3, "Group3"},
		{ModificationFlagParam, "Param"},
	}

	var parts []string
	for _, f := range flags {
		if mf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

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
