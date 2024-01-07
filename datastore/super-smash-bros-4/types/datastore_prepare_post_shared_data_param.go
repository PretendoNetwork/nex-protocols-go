// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePreparePostSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePreparePostSharedDataParam struct {
	types.Structure
	DataType   *types.PrimitiveU8
	Region     *types.PrimitiveU8
	Attribute1 *types.PrimitiveU8
	Attribute2 *types.PrimitiveU8
	Fighter    *types.Buffer
	Size       *types.PrimitiveU32
	Comment    *types.String
	MetaBinary *types.QBuffer
	ExtraData  *types.List[*types.String]
}

// ExtractFrom extracts the DataStorePreparePostSharedDataParam from the given readable
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePreparePostSharedDataParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePreparePostSharedDataParam header. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.DataType. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Region. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute1. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute2. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Fighter. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Size. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.Comment.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Comment. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.MetaBinary. %s", err.Error())
	}

	err = dataStorePreparePostSharedDataParam.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.ExtraData. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePreparePostSharedDataParam to the given writable
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePreparePostSharedDataParam.DataType.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Region.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Attribute1.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Attribute2.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Fighter.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Size.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.Comment.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.MetaBinary.WriteTo(contentWritable)
	dataStorePreparePostSharedDataParam.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePreparePostSharedDataParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePreparePostSharedDataParam
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Copy() types.RVType {
	copied := NewDataStorePreparePostSharedDataParam()

	copied.StructureVersion = dataStorePreparePostSharedDataParam.StructureVersion

	copied.DataType = dataStorePreparePostSharedDataParam.DataType.Copy().(*types.PrimitiveU8)
	copied.Region = dataStorePreparePostSharedDataParam.Region.Copy().(*types.PrimitiveU8)
	copied.Attribute1 = dataStorePreparePostSharedDataParam.Attribute1.Copy().(*types.PrimitiveU8)
	copied.Attribute2 = dataStorePreparePostSharedDataParam.Attribute2.Copy().(*types.PrimitiveU8)
	copied.Fighter = dataStorePreparePostSharedDataParam.Fighter.Copy().(*types.Buffer)
	copied.Size = dataStorePreparePostSharedDataParam.Size.Copy().(*types.PrimitiveU32)
	copied.Comment = dataStorePreparePostSharedDataParam.Comment.Copy().(*types.String)
	copied.MetaBinary = dataStorePreparePostSharedDataParam.MetaBinary.Copy().(*types.QBuffer)
	copied.ExtraData = dataStorePreparePostSharedDataParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostSharedDataParam)

	if dataStorePreparePostSharedDataParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePreparePostSharedDataParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Region.Equals(other.Region) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Attribute2.Equals(other.Attribute2) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Size.Equals(other.Size) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.Comment.Equals(other.Comment) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostSharedDataParam.ExtraData.Equals(other.ExtraData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) String() string {
	return dataStorePreparePostSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.DataType))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Fighter))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Size))
	b.WriteString(fmt.Sprintf("%sComment: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.Comment))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStorePreparePostSharedDataParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sExtraData: %s\n", indentationValues, dataStorePreparePostSharedDataParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostSharedDataParam returns a new DataStorePreparePostSharedDataParam
func NewDataStorePreparePostSharedDataParam() *DataStorePreparePostSharedDataParam {
	dataStorePreparePostSharedDataParam := &DataStorePreparePostSharedDataParam{
		DataType: types.NewPrimitiveU8(0),
		Region: types.NewPrimitiveU8(0),
		Attribute1: types.NewPrimitiveU8(0),
		Attribute2: types.NewPrimitiveU8(0),
		Fighter: types.NewBuffer(nil),
		Size: types.NewPrimitiveU32(0),
		Comment: types.NewString(""),
		MetaBinary: types.NewQBuffer(nil),
		ExtraData: types.NewList[*types.String](),
	}

	dataStorePreparePostSharedDataParam.ExtraData.Type = types.NewString("")

	return dataStorePreparePostSharedDataParam
}
