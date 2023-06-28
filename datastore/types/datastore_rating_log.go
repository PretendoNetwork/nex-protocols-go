package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreRatingLog struct {
	nex.Structure
	IsRated            bool
	Pid                uint32
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

	dataStoreRatingLog.Pid, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.Pid. %s", err.Error())
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
	stream.WriteUInt32LE(dataStoreRatingLog.Pid)
	stream.WriteInt32LE(dataStoreRatingLog.RatingValue)
	stream.WriteDateTime(dataStoreRatingLog.LockExpirationTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRatingLog
func (dataStoreRatingLog *DataStoreRatingLog) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingLog()

	copied.IsRated = dataStoreRatingLog.IsRated
	copied.Pid = dataStoreRatingLog.Pid
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

	if dataStoreRatingLog.Pid != other.Pid {
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

// NewDataStoreRatingLog returns a new DataStoreRatingLog
func NewDataStoreRatingLog() *DataStoreRatingLog {
	return &DataStoreRatingLog{}
}
