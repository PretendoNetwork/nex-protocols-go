package datastore_super_mario_maker_types

import "github.com/PretendoNetwork/nex-go"

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

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	return &DataStoreGetCourseRecordResult{}
}
