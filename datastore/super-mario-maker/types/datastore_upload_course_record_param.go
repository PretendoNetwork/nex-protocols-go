package datastore_super_mario_maker_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

// NewDataStoreUploadCourseRecordParam returns a new DataStoreUploadCourseRecordParam
func NewDataStoreUploadCourseRecordParam() *DataStoreUploadCourseRecordParam {
	return &DataStoreUploadCourseRecordParam{}
}
