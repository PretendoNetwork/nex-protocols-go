// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreUploadCourseRecordParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreUploadCourseRecordParam struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
	Score  *types.PrimitiveS32
}

// ExtractFrom extracts the DataStoreUploadCourseRecordParam from the given readable
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreUploadCourseRecordParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreUploadCourseRecordParam header. %s", err.Error())
	}

	err = dataStoreUploadCourseRecordParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.DataID from stream. %s", err.Error())
	}

	err = dataStoreUploadCourseRecordParam.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Slot from stream. %s", err.Error())
	}

	err = dataStoreUploadCourseRecordParam.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Score from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreUploadCourseRecordParam to the given writable
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreUploadCourseRecordParam.DataID.WriteTo(contentWritable)
	dataStoreUploadCourseRecordParam.Slot.WriteTo(contentWritable)
	dataStoreUploadCourseRecordParam.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreUploadCourseRecordParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreUploadCourseRecordParam
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Copy() types.RVType {
	copied := NewDataStoreUploadCourseRecordParam()

	copied.StructureVersion = dataStoreUploadCourseRecordParam.StructureVersion

	copied.DataID = dataStoreUploadCourseRecordParam.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dataStoreUploadCourseRecordParam.Slot.Copy().(*types.PrimitiveU8)
	copied.Score = dataStoreUploadCourseRecordParam.Score.Copy().(*types.PrimitiveS32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreUploadCourseRecordParam); !ok {
		return false
	}

	other := o.(*DataStoreUploadCourseRecordParam)

	if dataStoreUploadCourseRecordParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreUploadCourseRecordParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreUploadCourseRecordParam.Slot.Equals(other.Slot) {
		return false
	}

	if !dataStoreUploadCourseRecordParam.Score.Equals(other.Score) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) String() string {
	return dataStoreUploadCourseRecordParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreUploadCourseRecordParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreUploadCourseRecordParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreUploadCourseRecordParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreUploadCourseRecordParam.Slot))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStoreUploadCourseRecordParam.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreUploadCourseRecordParam returns a new DataStoreUploadCourseRecordParam
func NewDataStoreUploadCourseRecordParam() *DataStoreUploadCourseRecordParam {
	return &DataStoreUploadCourseRecordParam{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
		Score:  types.NewPrimitiveS32(0),
	}
}
