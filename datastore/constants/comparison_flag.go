package constants

// ComparisonFlag indicates the flags set on comparisonFlag of DataStoreChangeMetaCompareParam.
// These flags tell the server what values to use when comparing
// objects during search
type ComparisonFlag uint16

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
