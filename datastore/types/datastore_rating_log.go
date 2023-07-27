// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingLog is a data structure used by the DataStore protocol
type DataStoreRatingLog struct {
	nex.Structure
	IsRated            bool
	PID                uint32
	RatingValue        int32
	LockExpirationTime *nex.DateTime
}

// ExtractFromStream extracts a DataStoreRatingLog structure from a stream
func (dataStoreRatingLog *DataStoreRatingLog) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingLog.IsRated, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.IsRated. %s", err.Error())
	}

	dataStoreRatingLog.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.PID. %s", err.Error())
	}

	dataStoreRatingLog.RatingValue, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.RatingValue. %s", err.Error())
	}

	dataStoreRatingLog.LockExpirationTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.LockExpirationTime. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreRatingLog and returns a byte array
func (dataStoreRatingLog *DataStoreRatingLog) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteBool(dataStoreRatingLog.IsRated)
	stream.WriteUInt32LE(dataStoreRatingLog.PID)
	stream.WriteInt32LE(dataStoreRatingLog.RatingValue)
	stream.WriteDateTime(dataStoreRatingLog.LockExpirationTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRatingLog
func (dataStoreRatingLog *DataStoreRatingLog) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingLog()

	copied.IsRated = dataStoreRatingLog.IsRated
	copied.PID = dataStoreRatingLog.PID
	copied.RatingValue = dataStoreRatingLog.RatingValue
	copied.LockExpirationTime = dataStoreRatingLog.LockExpirationTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingLog *DataStoreRatingLog) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingLog)

	if dataStoreRatingLog.IsRated != other.IsRated {
		return false
	}

	if dataStoreRatingLog.PID != other.PID {
		return false
	}

	if dataStoreRatingLog.RatingValue != other.RatingValue {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRatingLog.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sIsRated: %t,\n", indentationValues, dataStoreRatingLog.IsRated))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, dataStoreRatingLog.PID))
	b.WriteString(fmt.Sprintf("%sRatingValue: %d,\n", indentationValues, dataStoreRatingLog.RatingValue))

	if dataStoreRatingLog.LockExpirationTime != nil {
		b.WriteString(fmt.Sprintf("%sLockExpirationTime: %s\n", indentationValues, dataStoreRatingLog.LockExpirationTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLockExpirationTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingLog returns a new DataStoreRatingLog
func NewDataStoreRatingLog() *DataStoreRatingLog {
	return &DataStoreRatingLog{}
}
