// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetCourseRecordResult holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCourseRecordResult struct {
	types.Structure
	DataID      *types.PrimitiveU64
	Slot        *types.PrimitiveU8
	FirstPID    *types.PID
	BestPID     *types.PID
	BestScore   *types.PrimitiveS32
	CreatedTime *types.DateTime
	UpdatedTime *types.DateTime
}

// ExtractFrom extracts the DataStoreGetCourseRecordResult from the given readable
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetCourseRecordResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetCourseRecordResult header. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.DataID from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.Slot from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.FirstPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.FirstPID from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.BestPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestPID from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.BestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestScore from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.CreatedTime from stream. %s", err.Error())
	}

	err = dataStoreGetCourseRecordResult.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.UpdatedTime from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetCourseRecordResult to the given writable
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetCourseRecordResult.DataID.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.Slot.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.FirstPID.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.BestPID.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.BestScore.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.CreatedTime.WriteTo(contentWritable)
	dataStoreGetCourseRecordResult.UpdatedTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetCourseRecordResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetCourseRecordResult
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Copy() types.RVType {
	copied := NewDataStoreGetCourseRecordResult()

	copied.StructureVersion = dataStoreGetCourseRecordResult.StructureVersion

	copied.DataID = dataStoreGetCourseRecordResult.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dataStoreGetCourseRecordResult.Slot.Copy().(*types.PrimitiveU8)
	copied.FirstPID = dataStoreGetCourseRecordResult.FirstPID.Copy().(*types.PID)
	copied.BestPID = dataStoreGetCourseRecordResult.BestPID.Copy().(*types.PID)
	copied.BestScore = dataStoreGetCourseRecordResult.BestScore.Copy().(*types.PrimitiveS32)
	copied.CreatedTime = dataStoreGetCourseRecordResult.CreatedTime.Copy().(*types.DateTime)
	copied.UpdatedTime = dataStoreGetCourseRecordResult.UpdatedTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCourseRecordResult); !ok {
		return false
	}

	other := o.(*DataStoreGetCourseRecordResult)

	if dataStoreGetCourseRecordResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetCourseRecordResult.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreGetCourseRecordResult.Slot.Equals(other.Slot) {
		return false
	}

	if !dataStoreGetCourseRecordResult.FirstPID.Equals(other.FirstPID) {
		return false
	}

	if !dataStoreGetCourseRecordResult.BestPID.Equals(other.BestPID) {
		return false
	}

	if !dataStoreGetCourseRecordResult.BestScore.Equals(other.BestScore) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetCourseRecordResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreGetCourseRecordResult.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreGetCourseRecordResult.Slot))
	b.WriteString(fmt.Sprintf("%sFirstPID: %s,\n", indentationValues, dataStoreGetCourseRecordResult.FirstPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestPID: %s,\n", indentationValues, dataStoreGetCourseRecordResult.BestPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestScore: %s,\n", indentationValues, dataStoreGetCourseRecordResult.BestScore))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s\n", indentationValues, dataStoreGetCourseRecordResult.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, dataStoreGetCourseRecordResult.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	return &DataStoreGetCourseRecordResult{
		DataID:      types.NewPrimitiveU64(0),
		Slot:        types.NewPrimitiveU8(0),
		FirstPID:    types.NewPID(0),
		BestPID:     types.NewPID(0),
		BestScore:   types.NewPrimitiveS32(0),
		CreatedTime: types.NewDateTime(0),
		UpdatedTime: types.NewDateTime(0),
	}
}
