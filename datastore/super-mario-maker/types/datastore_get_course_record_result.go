// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCourseRecordResult holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCourseRecordResult struct {
	nex.Structure
	DataID      uint64
	Slot        uint8
	FirstPID    *nex.PID
	BestPID     *nex.PID
	BestScore   int32
	CreatedTime *nex.DateTime
	UpdatedTime *nex.DateTime
}

// ExtractFromStream extracts a DataStoreGetCourseRecordResult structure from a stream
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetCourseRecordResult.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.DataID from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.Slot from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.FirstPID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.FirstPID from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.BestPID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestPID from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.BestScore, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestScore from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.CreatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.CreatedTime from stream. %s", err.Error())
	}

	dataStoreGetCourseRecordResult.UpdatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.UpdatedTime from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetCourseRecordResult and returns a byte array
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.DataID)
	stream.WriteUInt8(dataStoreGetCourseRecordResult.Slot)
	stream.WritePID(dataStoreGetCourseRecordResult.FirstPID)
	stream.WritePID(dataStoreGetCourseRecordResult.BestPID)
	stream.WriteInt32LE(dataStoreGetCourseRecordResult.BestScore)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.CreatedTime)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.UpdatedTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetCourseRecordResult
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordResult()

	copied.SetStructureVersion(dataStoreGetCourseRecordResult.StructureVersion())

	copied.DataID = dataStoreGetCourseRecordResult.DataID
	copied.Slot = dataStoreGetCourseRecordResult.Slot
	copied.FirstPID = dataStoreGetCourseRecordResult.FirstPID.Copy()
	copied.BestPID = dataStoreGetCourseRecordResult.BestPID.Copy()
	copied.BestScore = dataStoreGetCourseRecordResult.BestScore
	copied.CreatedTime = dataStoreGetCourseRecordResult.CreatedTime.Copy()
	copied.UpdatedTime = dataStoreGetCourseRecordResult.UpdatedTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordResult)

	if dataStoreGetCourseRecordResult.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreGetCourseRecordResult.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordResult.Slot != other.Slot {
		return false
	}

	if !dataStoreGetCourseRecordResult.FirstPID.Equals(other.FirstPID) {
		return false
	}

	if !dataStoreGetCourseRecordResult.BestPID.Equals(other.BestPID) {
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
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreGetCourseRecordResult.Slot))
	b.WriteString(fmt.Sprintf("%sFirstPID: %s,\n", indentationValues, dataStoreGetCourseRecordResult.FirstPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestPID: %s,\n", indentationValues, dataStoreGetCourseRecordResult.BestPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestScore: %d,\n", indentationValues, dataStoreGetCourseRecordResult.BestScore))

	if dataStoreGetCourseRecordResult.CreatedTime != nil {
		b.WriteString(fmt.Sprintf("%sCreatedTime: %s\n", indentationValues, dataStoreGetCourseRecordResult.CreatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedTime: nil\n", indentationValues))
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
	return &DataStoreGetCourseRecordResult{
		DataID:      0,
		Slot:        0,
		FirstPID:    nex.NewPID[uint32](0),
		BestPID:     nex.NewPID[uint32](0),
		BestScore:   0,
		CreatedTime: nex.NewDateTime(0),
		UpdatedTime: nex.NewDateTime(0),
	}
}
