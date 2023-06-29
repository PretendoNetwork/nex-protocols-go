package datastore_super_mario_maker_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCourseRecordResult is used to send data about a courses world record
type DataStoreGetCourseRecordResult struct {
	nex.Structure
	DataID      uint64
	Slot        uint8
	FirstPID    uint32
	BestPID     uint32
	BestScore   int32
	CreatedTime *nex.DateTime
	UpdatedTime *nex.DateTime
}

// Bytes encodes the DataStoreGetCourseRecordResult and returns a byte array
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.DataID)
	stream.WriteUInt8(dataStoreGetCourseRecordResult.Slot)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.FirstPID)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.BestPID)
	stream.WriteInt32LE(dataStoreGetCourseRecordResult.BestScore)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.CreatedTime)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.UpdatedTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetCourseRecordResult
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordResult()

	copied.DataID = dataStoreGetCourseRecordResult.DataID
	copied.Slot = dataStoreGetCourseRecordResult.Slot
	copied.FirstPID = dataStoreGetCourseRecordResult.FirstPID
	copied.BestPID = dataStoreGetCourseRecordResult.BestPID
	copied.BestScore = dataStoreGetCourseRecordResult.BestScore
	copied.CreatedTime = dataStoreGetCourseRecordResult.CreatedTime.Copy()
	copied.UpdatedTime = dataStoreGetCourseRecordResult.UpdatedTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordResult)

	if dataStoreGetCourseRecordResult.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordResult.Slot != other.Slot {
		return false
	}

	if dataStoreGetCourseRecordResult.FirstPID != other.FirstPID {
		return false
	}

	if dataStoreGetCourseRecordResult.BestPID != other.BestPID {
		return false
	}

	if dataStoreGetCourseRecordResult.BestScore != other.BestScore {
		return false
	}

	if !dataStoreGetCourseRecordResult.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreGetCourseRecordResult.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) String() string {
	return dataStoreGetCourseRecordResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCourseRecordResult{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetCourseRecordResult.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreGetCourseRecordResult.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %d\n,", indentationValues, dataStoreGetCourseRecordResult.Slot))
	b.WriteString(fmt.Sprintf("%sFirstPID: %d,\n", indentationValues, dataStoreGetCourseRecordResult.FirstPID))
	b.WriteString(fmt.Sprintf("%sBestPID: %d,\n", indentationValues, dataStoreGetCourseRecordResult.BestPID))
	b.WriteString(fmt.Sprintf("%sBestScore: %d,\n", indentationValues, dataStoreGetCourseRecordResult.BestScore))

	if dataStoreGetCourseRecordResult.CreatedTime != nil {
		b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dataStoreGetCourseRecordResult.CreatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedTime: nil,\n", indentationValues))
	}

	if dataStoreGetCourseRecordResult.UpdatedTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, dataStoreGetCourseRecordResult.UpdatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	return &DataStoreGetCourseRecordResult{}
}
