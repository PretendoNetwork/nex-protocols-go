package datastore_super_mario_maker_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCourseRecordParam is sent in the GetMetasWithCourseRecord method
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
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.DataID. %s", err.Error())
	}

	dataStoreGetCourseRecordParam.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCourseRecordParam
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordParam()

	copied.DataID = dataStoreGetCourseRecordParam.DataID
	copied.Slot = dataStoreGetCourseRecordParam.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordParam)

	if dataStoreGetCourseRecordParam.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordParam.Slot != other.Slot {
		return false
	}

	return true
}

// NewDataStoreGetCourseRecordParamreturns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	return &DataStoreGetCourseRecordParam{}
}
