// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCourseRecordParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCourseRecordParam struct {
	nex.Structure
	DataID uint64
	Slot   uint8
}

// ExtractFromStream extracts a DataStoreGetCourseRecordParam structure from a stream
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetCourseRecordParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.DataID from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordParam.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.Slot from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetCourseRecordParam and returns a byte array
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetCourseRecordParam.DataID)
	stream.WriteUInt8(dataStoreGetCourseRecordParam.Slot)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetCourseRecordParam
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordParam()

	copied.SetStructureVersion(dataStoreGetCourseRecordParam.StructureVersion())

	copied.DataID = dataStoreGetCourseRecordParam.DataID
	copied.Slot = dataStoreGetCourseRecordParam.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordParam)

	if dataStoreGetCourseRecordParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreGetCourseRecordParam.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordParam.Slot != other.Slot {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetCourseRecordParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreGetCourseRecordParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreGetCourseRecordParam.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordParam returns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	return &DataStoreGetCourseRecordParam{
		DataID: 0,
		Slot:   0,
	}
}
