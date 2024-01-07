// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetMetaByOwnerIDParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetMetaByOwnerIDParam struct {
	types.Structure
	OwnerIDs     *types.List[*types.PrimitiveU32]
	DataTypes    *types.List[*types.PrimitiveU16]
	ResultOption *types.PrimitiveU8
	ResultRange  *types.ResultRange
}

// ExtractFrom extracts the DataStoreGetMetaByOwnerIDParam from the given readable
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetMetaByOwnerIDParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetMetaByOwnerIDParam header. %s", err.Error())
	}

	err = dataStoreGetMetaByOwnerIDParam.OwnerIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.OwnerIDs from stream. %s", err.Error())
	}

	err = dataStoreGetMetaByOwnerIDParam.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.DataTypes from stream. %s", err.Error())
	}

	err = dataStoreGetMetaByOwnerIDParam.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption from stream. %s", err.Error())
	}

	err = dataStoreGetMetaByOwnerIDParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultRange from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetMetaByOwnerIDParam to the given writable
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetMetaByOwnerIDParam.OwnerIDs.WriteTo(contentWritable)
	dataStoreGetMetaByOwnerIDParam.DataTypes.WriteTo(contentWritable)
	dataStoreGetMetaByOwnerIDParam.ResultOption.WriteTo(contentWritable)
	dataStoreGetMetaByOwnerIDParam.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetMetaByOwnerIDParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetMetaByOwnerIDParam
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByOwnerIDParam()

	copied.StructureVersion = dataStoreGetMetaByOwnerIDParam.StructureVersion

	copied.OwnerIDs = dataStoreGetMetaByOwnerIDParam.OwnerIDs.Copy().(*types.List[*types.PrimitiveU32])
	copied.DataTypes = dataStoreGetMetaByOwnerIDParam.DataTypes.Copy().(*types.List[*types.PrimitiveU16])
	copied.ResultOption = dataStoreGetMetaByOwnerIDParam.ResultOption.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dataStoreGetMetaByOwnerIDParam.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByOwnerIDParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByOwnerIDParam)

	if dataStoreGetMetaByOwnerIDParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.OwnerIDs.Equals(other.OwnerIDs) {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.DataTypes.Equals(other.DataTypes) {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) String() string {
	return dataStoreGetMetaByOwnerIDParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByOwnerIDParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %s,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreGetMetaByOwnerIDParam.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	dataStoreGetMetaByOwnerIDParam := &DataStoreGetMetaByOwnerIDParam{
		OwnerIDs:     types.NewList[*types.PrimitiveU32](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	dataStoreGetMetaByOwnerIDParam.OwnerIDs.Type = types.NewPrimitiveU32(0)
	dataStoreGetMetaByOwnerIDParam.DataTypes.Type = types.NewPrimitiveU16(0)

	return dataStoreGetMetaByOwnerIDParam
}
