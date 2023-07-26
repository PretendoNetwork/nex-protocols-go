// Package types implements all the types used by the DataStore Super Mario Maker protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreUploadCourseRecordParam is a data structure used by the DataStore Super Mario Maker protocol
type DataStoreUploadCourseRecordParam struct {
	nex.Structure
	DataID uint64
	Slot   uint8
	Score  int32
}

// ExtractFromStream extracts a DataStoreUploadCourseRecordParam structure from a stream
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreUploadCourseRecordParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.DataID. %s", err.Error())
	}

	dataStoreUploadCourseRecordParam.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Slot. %s", err.Error())
	}

	dataStoreUploadCourseRecordParam.Score, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreUploadCourseRecordParam
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Copy() nex.StructureInterface {
	copied := NewDataStoreUploadCourseRecordParam()

	copied.DataID = dataStoreUploadCourseRecordParam.DataID
	copied.Slot = dataStoreUploadCourseRecordParam.Slot
	copied.Score = dataStoreUploadCourseRecordParam.Score

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreUploadCourseRecordParam)

	if dataStoreUploadCourseRecordParam.DataID != other.DataID {
		return false
	}

	if dataStoreUploadCourseRecordParam.Slot != other.Slot {
		return false
	}

	if dataStoreUploadCourseRecordParam.Score != other.Score {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreUploadCourseRecordParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreUploadCourseRecordParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreUploadCourseRecordParam.Slot))
	b.WriteString(fmt.Sprintf("%sScore: %d\n", indentationValues, dataStoreUploadCourseRecordParam.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreUploadCourseRecordParam returns a new DataStoreUploadCourseRecordParam
func NewDataStoreUploadCourseRecordParam() *DataStoreUploadCourseRecordParam {
	return &DataStoreUploadCourseRecordParam{}
}
