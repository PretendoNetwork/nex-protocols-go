// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingLog is a data structure used by the DataStore protocol
type DataStoreRatingLog struct {
	types.Structure
	IsRated            *types.PrimitiveBool
	PID                *types.PID
	RatingValue        *types.PrimitiveS32
	LockExpirationTime *types.DateTime
}

// ExtractFrom extracts the DataStoreRatingLog from the given readable
func (dataStoreRatingLog *DataStoreRatingLog) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingLog.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingLog header. %s", err.Error())
	}

	err = dataStoreRatingLog.IsRated.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.IsRated. %s", err.Error())
	}

	err = dataStoreRatingLog.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.PID. %s", err.Error())
	}

	err = dataStoreRatingLog.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.RatingValue. %s", err.Error())
	}

	err = dataStoreRatingLog.LockExpirationTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.LockExpirationTime. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreRatingLog to the given writable
func (dataStoreRatingLog *DataStoreRatingLog) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingLog.IsRated.WriteTo(contentWritable)
	dataStoreRatingLog.PID.WriteTo(contentWritable)
	dataStoreRatingLog.RatingValue.WriteTo(contentWritable)
	dataStoreRatingLog.LockExpirationTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingLog.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreRatingLog
func (dataStoreRatingLog *DataStoreRatingLog) Copy() types.RVType {
	copied := NewDataStoreRatingLog()

	copied.StructureVersion = dataStoreRatingLog.StructureVersion

	copied.IsRated = dataStoreRatingLog.IsRated.Copy().(*types.PrimitiveBool)
	copied.PID = dataStoreRatingLog.PID.Copy().(*types.PID)
	copied.RatingValue = dataStoreRatingLog.RatingValue.Copy().(*types.PrimitiveS32)
	copied.LockExpirationTime = dataStoreRatingLog.LockExpirationTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingLog *DataStoreRatingLog) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingLog); !ok {
		return false
	}

	other := o.(*DataStoreRatingLog)

	if dataStoreRatingLog.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingLog.IsRated.Equals(other.IsRated) {
		return false
	}

	if !dataStoreRatingLog.PID.Equals(other.PID) {
		return false
	}

	if !dataStoreRatingLog.RatingValue.Equals(other.RatingValue) {
		return false
	}

	if !dataStoreRatingLog.LockExpirationTime.Equals(other.LockExpirationTime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRatingLog *DataStoreRatingLog) String() string {
	return dataStoreRatingLog.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingLog *DataStoreRatingLog) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingLog{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingLog.StructureVersion))
	b.WriteString(fmt.Sprintf("%sIsRated: %s,\n", indentationValues, dataStoreRatingLog.IsRated))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, dataStoreRatingLog.PID))
	b.WriteString(fmt.Sprintf("%sRatingValue: %s,\n", indentationValues, dataStoreRatingLog.RatingValue))
	b.WriteString(fmt.Sprintf("%sLockExpirationTime: %s\n", indentationValues, dataStoreRatingLog.LockExpirationTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingLog returns a new DataStoreRatingLog
func NewDataStoreRatingLog() *DataStoreRatingLog {
	return &DataStoreRatingLog{
		IsRated:            types.NewPrimitiveBool(false),
		PID:                types.NewPID(0),
		RatingValue:        types.NewPrimitiveS32(0),
		LockExpirationTime: types.NewDateTime(0),
	}
}
