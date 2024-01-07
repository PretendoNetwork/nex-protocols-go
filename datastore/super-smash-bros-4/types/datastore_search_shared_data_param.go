// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSearchSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreSearchSharedDataParam struct {
	types.Structure
	DataType    *types.PrimitiveU8
	Owner       *types.PrimitiveU32
	Region      *types.PrimitiveU8
	Attribute1  *types.PrimitiveU8
	Attribute2  *types.PrimitiveU8
	Fighter     *types.PrimitiveU8
	ResultRange *types.ResultRange
}

// ExtractFrom extracts the DataStoreSearchSharedDataParam from the given readable
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreSearchSharedDataParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSearchSharedDataParam header. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.DataType. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.Owner.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Owner. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Region. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute1. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute2. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Fighter. %s", err.Error())
	}

	err = dataStoreSearchSharedDataParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.ResultRange. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreSearchSharedDataParam to the given writable
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreSearchSharedDataParam.DataType.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.Owner.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.Region.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.Attribute1.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.Attribute2.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.Fighter.WriteTo(contentWritable)
	dataStoreSearchSharedDataParam.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreSearchSharedDataParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreSearchSharedDataParam
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Copy() types.RVType {
	copied := NewDataStoreSearchSharedDataParam()

	copied.StructureVersion = dataStoreSearchSharedDataParam.StructureVersion

	copied.DataType = dataStoreSearchSharedDataParam.DataType
	copied.Owner = dataStoreSearchSharedDataParam.Owner
	copied.Region = dataStoreSearchSharedDataParam.Region
	copied.Attribute1 = dataStoreSearchSharedDataParam.Attribute1
	copied.Attribute2 = dataStoreSearchSharedDataParam.Attribute2
	copied.Fighter = dataStoreSearchSharedDataParam.Fighter
	copied.ResultRange = dataStoreSearchSharedDataParam.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchSharedDataParam)

	if dataStoreSearchSharedDataParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreSearchSharedDataParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreSearchSharedDataParam.Owner.Equals(other.Owner) {
		return false
	}

	if !dataStoreSearchSharedDataParam.Region.Equals(other.Region) {
		return false
	}

	if !dataStoreSearchSharedDataParam.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !dataStoreSearchSharedDataParam.Attribute2.Equals(other.Attribute2) {
		return false
	}

	if !dataStoreSearchSharedDataParam.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dataStoreSearchSharedDataParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) String() string {
	return dataStoreSearchSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreSearchSharedDataParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSearchSharedDataParam.DataType))
	b.WriteString(fmt.Sprintf("%sOwner: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Owner))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Fighter))

	if dataStoreSearchSharedDataParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreSearchSharedDataParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchSharedDataParam returns a new DataStoreSearchSharedDataParam
func NewDataStoreSearchSharedDataParam() *DataStoreSearchSharedDataParam {
	return &DataStoreSearchSharedDataParam{}
}
