// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetCourseRecordParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCourseRecordParam struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
}

// ExtractFrom extracts the DataStoreGetCourseRecordParam from the given readable
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetCourseRecordParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetCourseRecordParam header. %s", err.Error())
	}

	err = dataStoreGetCourseRecordParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.DataID from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordParam.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.Slot from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetCourseRecordParam to the given writable
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetCourseRecordParam.DataID.WriteTo(contentWritable)
	dataStoreGetCourseRecordParam.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetCourseRecordParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetCourseRecordParam
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Copy() types.RVType {
	copied := NewDataStoreGetCourseRecordParam()

	copied.StructureVersion = dataStoreGetCourseRecordParam.StructureVersion

	copied.DataID = dataStoreGetCourseRecordParam.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dataStoreGetCourseRecordParam.Slot.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCourseRecordParam); !ok {
		return false
	}

	other := o.(*DataStoreGetCourseRecordParam)

	if dataStoreGetCourseRecordParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetCourseRecordParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreGetCourseRecordParam.Slot.Equals(other.Slot) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) String() string {
	return dataStoreGetCourseRecordParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCourseRecordParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetCourseRecordParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreGetCourseRecordParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreGetCourseRecordParam.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordParam returns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	return &DataStoreGetCourseRecordParam{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
	}
}
