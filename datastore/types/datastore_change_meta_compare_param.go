// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreChangeMetaCompareParam is sent in the ChangeMeta method
type DataStoreChangeMetaCompareParam struct {
	types.Structure
	ComparisonFlag *types.PrimitiveU32
	Name           *types.String
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         *types.PrimitiveU16
	MetaBinary     *types.QBuffer
	Tags           *types.List[*types.String]
	ReferredCnt    *types.PrimitiveU32
	DataType       *types.PrimitiveU16
	Status         *types.PrimitiveU8
}

// WriteTo writes the DataStoreChangeMetaCompareParam to the given writable
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreChangeMetaCompareParam.ComparisonFlag.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.Name.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.Permission.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.DelPermission.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.Period.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.MetaBinary.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.Tags.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.ReferredCnt.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.DataType.WriteTo(contentWritable)
	dataStoreChangeMetaCompareParam.Status.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreChangeMetaCompareParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaCompareParam from the given readable
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreChangeMetaCompareParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreChangeMetaCompareParam header. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.ComparisonFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ComparisonFlag. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Name. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Permission. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DelPermission. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Period. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.MetaBinary. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Tags. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ReferredCnt. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DataType. %s", err.Error())
	}

	err = dataStoreChangeMetaCompareParam.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Status. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaCompareParam
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaCompareParam()

	copied.StructureVersion = dataStoreChangeMetaCompareParam.StructureVersion

	copied.ComparisonFlag = dataStoreChangeMetaCompareParam.ComparisonFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dataStoreChangeMetaCompareParam.Name.Copy().(*types.String)
	copied.Permission = dataStoreChangeMetaCompareParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaCompareParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaCompareParam.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dataStoreChangeMetaCompareParam.MetaBinary.Copy().(*types.QBuffer)

	copied.Tags = dataStoreChangeMetaCompareParam.Tags.Copy().(*types.List[*types.String])

	copied.ReferredCnt = dataStoreChangeMetaCompareParam.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.DataType = dataStoreChangeMetaCompareParam.DataType.Copy().(*types.PrimitiveU16)
	copied.Status = dataStoreChangeMetaCompareParam.Status.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaCompareParam); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaCompareParam)

	if dataStoreChangeMetaCompareParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreChangeMetaCompareParam.ComparisonFlag.Equals(other.ComparisonFlag) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.Name.Equals(other.Name) {
		return false
	}

	if dataStoreChangeMetaCompareParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaCompareParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.Period.Equals(other.Period) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreChangeMetaCompareParam.Status.Equals(other.Status) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) String() string {
	return dataStoreChangeMetaCompareParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaCompareParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sComparisonFlag: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.ComparisonFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreChangeMetaCompareParam.Tags)) // TODO - Make this a nicer looking log
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %s\n", indentationValues, dataStoreChangeMetaCompareParam.Status))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	dataStoreChangeMetaCompareParam := &DataStoreChangeMetaCompareParam{
		ComparisonFlag: types.NewPrimitiveU32(0),
		Name:           types.NewString(""),
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         types.NewPrimitiveU16(0),
		MetaBinary:     types.NewQBuffer(nil),
		Tags:           types.NewList[*types.String](),
		ReferredCnt:    types.NewPrimitiveU32(0),
		DataType:       types.NewPrimitiveU16(0),
		Status:         types.NewPrimitiveU8(0),
	}

	dataStoreChangeMetaCompareParam.Tags.Type = types.NewString("")

	return dataStoreChangeMetaCompareParam
}
