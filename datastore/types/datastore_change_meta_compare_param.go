// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/datastore/constants"
)

// DataStoreChangeMetaCompareParam is a type within the DataStore protocol
type DataStoreChangeMetaCompareParam struct {
	types.Structure
	ComparisonFlag constants.ComparisonFlag
	Name           types.String
	Permission     DataStorePermission
	DelPermission  DataStorePermission
	Period         types.UInt16
	MetaBinary     types.QBuffer
	Tags           types.List[types.String]
	ReferredCnt    types.UInt32
	DataType       types.UInt16
	Status         constants.DataStatus
}

// WriteTo writes the DataStoreChangeMetaCompareParam to the given writable
func (dscmcp DataStoreChangeMetaCompareParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscmcp.ComparisonFlag.WriteTo(contentWritable)
	dscmcp.Name.WriteTo(contentWritable)
	dscmcp.Permission.WriteTo(contentWritable)
	dscmcp.DelPermission.WriteTo(contentWritable)
	dscmcp.Period.WriteTo(contentWritable)
	dscmcp.MetaBinary.WriteTo(contentWritable)
	dscmcp.Tags.WriteTo(contentWritable)
	dscmcp.ReferredCnt.WriteTo(contentWritable)
	dscmcp.DataType.WriteTo(contentWritable)
	dscmcp.Status.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscmcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaCompareParam from the given readable
func (dscmcp *DataStoreChangeMetaCompareParam) ExtractFrom(readable types.Readable) error {
	if err := dscmcp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam header. %s", err.Error())
	}

	if err := dscmcp.ComparisonFlag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.ComparisonFlag. %s", err.Error())
	}

	if err := dscmcp.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.Name. %s", err.Error())
	}

	if err := dscmcp.Permission.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.Permission. %s", err.Error())
	}

	if err := dscmcp.DelPermission.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.DelPermission. %s", err.Error())
	}

	if err := dscmcp.Period.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.Period. %s", err.Error())
	}

	if err := dscmcp.MetaBinary.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.MetaBinary. %s", err.Error())
	}

	if err := dscmcp.Tags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.Tags. %s", err.Error())
	}

	if err := dscmcp.ReferredCnt.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.ReferredCnt. %s", err.Error())
	}

	if err := dscmcp.DataType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.DataType. %s", err.Error())
	}

	if err := dscmcp.Status.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DataStoreChangeMetaCompareParam.Status. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaCompareParam
func (dscmcp DataStoreChangeMetaCompareParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaCompareParam()

	copied.StructureVersion = dscmcp.StructureVersion
	copied.ComparisonFlag = dscmcp.ComparisonFlag
	copied.Name = dscmcp.Name.Copy().(types.String)
	copied.Permission = dscmcp.Permission.Copy().(DataStorePermission)
	copied.DelPermission = dscmcp.DelPermission.Copy().(DataStorePermission)
	copied.Period = dscmcp.Period.Copy().(types.UInt16)
	copied.MetaBinary = dscmcp.MetaBinary.Copy().(types.QBuffer)
	copied.Tags = dscmcp.Tags.Copy().(types.List[types.String])
	copied.ReferredCnt = dscmcp.ReferredCnt.Copy().(types.UInt32)
	copied.DataType = dscmcp.DataType.Copy().(types.UInt16)
	copied.Status = dscmcp.Status

	return copied
}

// Equals checks if the given DataStoreChangeMetaCompareParam contains the same data as the current DataStoreChangeMetaCompareParam
func (dscmcp DataStoreChangeMetaCompareParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreChangeMetaCompareParam); !ok {
		return false
	}

	other := o.(DataStoreChangeMetaCompareParam)

	if dscmcp.StructureVersion != other.StructureVersion {
		return false
	}

	if dscmcp.ComparisonFlag != other.ComparisonFlag {
		return false
	}

	if !dscmcp.Name.Equals(other.Name) {
		return false
	}

	if !dscmcp.Permission.Equals(other.Permission) {
		return false
	}

	if !dscmcp.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dscmcp.Period.Equals(other.Period) {
		return false
	}

	if !dscmcp.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dscmcp.Tags.Equals(other.Tags) {
		return false
	}

	if !dscmcp.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dscmcp.DataType.Equals(other.DataType) {
		return false
	}

	return dscmcp.Status == other.Status
}

// CopyRef copies the current value of the DataStoreChangeMetaCompareParam
// and returns a pointer to the new copy
func (dscmcp DataStoreChangeMetaCompareParam) CopyRef() types.RVTypePtr {
	copied := dscmcp.Copy().(DataStoreChangeMetaCompareParam)
	return &copied
}

// Deref takes a pointer to the DataStoreChangeMetaCompareParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dscmcp *DataStoreChangeMetaCompareParam) Deref() types.RVType {
	return *dscmcp
}

// String returns the string representation of the DataStoreChangeMetaCompareParam
func (dscmcp DataStoreChangeMetaCompareParam) String() string {
	return dscmcp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreChangeMetaCompareParam using the provided indentation level
func (dscmcp DataStoreChangeMetaCompareParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaCompareParam{\n")
	fmt.Fprintf(&b, "%sComparisonFlag: %s,\n", indentationValues, dscmcp.ComparisonFlag)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, dscmcp.Name)
	fmt.Fprintf(&b, "%sPermission: %s,\n", indentationValues, dscmcp.Permission.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sDelPermission: %s,\n", indentationValues, dscmcp.DelPermission.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sPeriod: %s,\n", indentationValues, dscmcp.Period)
	fmt.Fprintf(&b, "%sMetaBinary: %s,\n", indentationValues, dscmcp.MetaBinary)
	fmt.Fprintf(&b, "%sTags: %s,\n", indentationValues, dscmcp.Tags)
	fmt.Fprintf(&b, "%sReferredCnt: %s,\n", indentationValues, dscmcp.ReferredCnt)
	fmt.Fprintf(&b, "%sDataType: %s,\n", indentationValues, dscmcp.DataType)
	fmt.Fprintf(&b, "%sStatus: %s,\n", indentationValues, dscmcp.Status)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
func NewDataStoreChangeMetaCompareParam() DataStoreChangeMetaCompareParam {
	return DataStoreChangeMetaCompareParam{
		ComparisonFlag: constants.ComparisonFlagNone,
		Name:           types.NewString(""),
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         types.NewUInt16(0),
		MetaBinary:     types.NewQBuffer(nil),
		Tags:           types.NewList[types.String](),
		ReferredCnt:    types.NewUInt32(0),
		DataType:       types.NewUInt16(0),
		Status:         constants.DataStatusNone,
	}

}
