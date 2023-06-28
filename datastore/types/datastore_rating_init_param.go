package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingInitParam is sent in the PreparePostObject method
type DataStoreRatingInitParam struct {
	nex.Structure
	Flag           uint8
	InternalFlag   uint8
	LockType       uint8
	InitialValue   int64
	RangeMin       int32
	RangeMax       int32
	PeriodHour     int8
	PeriodDuration int16
}

// ExtractFromStream extracts a DataStoreRatingInitParam structure from a stream
func (dataStoreRatingInitParam *DataStoreRatingInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingInitParam.Flag, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.Flag. %s", err.Error())
	}

	dataStoreRatingInitParam.InternalFlag, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InternalFlag. %s", err.Error())
	}

	dataStoreRatingInitParam.LockType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.LockType. %s", err.Error())
	}

	dataStoreRatingInitParam.InitialValue, err = stream.ReadInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InitialValue. %s", err.Error())
	}

	dataStoreRatingInitParam.RangeMin, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMin. %s", err.Error())
	}

	dataStoreRatingInitParam.RangeMax, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMax. %s", err.Error())
	}

	dataStoreRatingInitParam.PeriodHour, err = stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodHour. %s", err.Error())
	}

	dataStoreRatingInitParam.PeriodDuration, err = stream.ReadInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodDuration. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParam
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInitParam()

	copied.Flag = dataStoreRatingInitParam.Flag
	copied.InternalFlag = dataStoreRatingInitParam.InternalFlag
	copied.LockType = dataStoreRatingInitParam.LockType
	copied.InitialValue = dataStoreRatingInitParam.InitialValue
	copied.RangeMin = dataStoreRatingInitParam.RangeMin
	copied.RangeMax = dataStoreRatingInitParam.RangeMax
	copied.PeriodHour = dataStoreRatingInitParam.PeriodHour
	copied.PeriodDuration = dataStoreRatingInitParam.PeriodDuration

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInitParam)

	if dataStoreRatingInitParam.Flag != other.Flag {
		return false
	}

	if dataStoreRatingInitParam.InternalFlag != other.InternalFlag {
		return false
	}

	if dataStoreRatingInitParam.LockType != other.LockType {
		return false
	}

	if dataStoreRatingInitParam.InitialValue != other.InitialValue {
		return false
	}

	if dataStoreRatingInitParam.RangeMin != other.RangeMin {
		return false
	}

	if dataStoreRatingInitParam.RangeMax != other.RangeMax {
		return false
	}

	if dataStoreRatingInitParam.PeriodHour != other.PeriodHour {
		return false
	}

	if dataStoreRatingInitParam.PeriodDuration != other.PeriodDuration {
		return false
	}

	return true
}

// NewDataStoreRatingInitParam returns a new DataStoreRatingInitParam
func NewDataStoreRatingInitParam() *DataStoreRatingInitParam {
	return &DataStoreRatingInitParam{}
}
