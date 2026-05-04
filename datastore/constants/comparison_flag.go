package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ComparisonFlag indicates the flags set on comparisonFlag of DataStoreChangeMetaCompareParam.
// These flags tell the server what values to use when comparing
// objects during search
type ComparisonFlag uint16

// WriteTo writes the ComparisonFlag to the given writable
func (cf ComparisonFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt16LE(uint16(cf))
}

// ExtractFrom extracts the ComparisonFlag value from the given readable
func (cf *ComparisonFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt16LE()
	if err != nil {
		return err
	}

	*cf = ComparisonFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (cf ComparisonFlag) HasFlag(flag ComparisonFlag) bool {
	return cf&flag == flag
}

// HasFlag checks if all given flags are set
func (cf ComparisonFlag) HasFlags(flags ...ComparisonFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if cf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the ComparisonFlag bitmask.
// Multiple flags are joined with "|", e.g. "Name|Period|Tags".
// Returns "None" if no flags are set, or "All" if all flags are set.
func (cf ComparisonFlag) String() string {
	if cf == ComparisonFlagNone {
		return "None"
	}
	if cf == ComparisonFlagAll {
		return "All"
	}

	flags := []struct {
		flag ComparisonFlag
		name string
	}{
		{ComparisonFlagName, "Name"},
		{ComparisonFlagAccessPermission, "AccessPermission"},
		{ComparisonFlagUpdatePermission, "UpdatePermission"},
		{ComparisonFlagPeriod, "Period"},
		{ComparisonFlagMetaBinary, "MetaBinary"},
		{ComparisonFlagTags, "Tags"},
		{ComparisonFlagDataType, "DataType"},
		{ComparisonFlagReferredCount, "ReferredCount"},
		{ComparisonFlagStatus, "Status"},
	}

	var parts []string
	for _, f := range flags {
		if cf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// ComparisonFlagNone means that no fields should be compared
	ComparisonFlagNone ComparisonFlag = 0x0

	// ComparisonFlagName means that the DataStoreChangeMetaCompareParam.name
	// field should be compared
	ComparisonFlagName ComparisonFlag = 0x1

	// ComparisonFlagAccessPermission means that the DataStoreChangeMetaCompareParam.permission
	// field should be compared
	ComparisonFlagAccessPermission ComparisonFlag = 0x2

	// ComparisonFlagUpdatePermission means that the DataStoreChangeMetaCompareParam.delPermission
	// field should be compared
	ComparisonFlagUpdatePermission ComparisonFlag = 0x4

	// ComparisonFlagPeriod means that the DataStoreChangeMetaCompareParam.period
	// field should be compared
	ComparisonFlagPeriod ComparisonFlag = 0x8

	// ComparisonFlagMetaBinary means that the DataStoreChangeMetaCompareParam.metaBinary
	// field should be compared
	ComparisonFlagMetaBinary ComparisonFlag = 0x10

	// ComparisonFlagTags means that the DataStoreChangeMetaCompareParam.tags
	// field should be compared
	ComparisonFlagTags ComparisonFlag = 0x20

	// ComparisonFlagDataType means that the DataStoreChangeMetaCompareParam.dataType
	// field should be compared
	ComparisonFlagDataType ComparisonFlag = 0x40

	// ComparisonFlagReferredCount means that the DataStoreChangeMetaCompareParam.referredCnt
	// field should be compared
	ComparisonFlagReferredCount ComparisonFlag = 0x80

	// ComparisonFlagStatus means that the DataStoreChangeMetaCompareParam.status
	// field should be compared
	ComparisonFlagStatus ComparisonFlag = 0x100

	// ComparisonFlagAll means that all fields should be compared.
	// Equivalent to setting each of the previous flags individually
	ComparisonFlagAll ComparisonFlag = 0xFFFF
)
