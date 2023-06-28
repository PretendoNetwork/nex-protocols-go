package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreGetReplayMetaParam struct {
	nex.Structure
	ReplayID uint64
	MetaType uint8
}

// ExtractFromStream extracts a DataStoreGetReplayMetaParam structure from a stream
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetReplayMetaParam.ReplayID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.ReplayID. %s", err.Error())
	}

	dataStoreGetReplayMetaParam.MetaType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.MetaType. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetReplayMetaParam and returns a byte array
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetReplayMetaParam.ReplayID)
	stream.WriteUInt8(dataStoreGetReplayMetaParam.MetaType)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetReplayMetaParam
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetReplayMetaParam()

	copied.ReplayID = dataStoreGetReplayMetaParam.ReplayID
	copied.MetaType = dataStoreGetReplayMetaParam.MetaType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetReplayMetaParam)

	if dataStoreGetReplayMetaParam.ReplayID != other.ReplayID {
		return false
	}

	if dataStoreGetReplayMetaParam.MetaType != other.MetaType {
		return false
	}

	return true
}

// NewDataStoreGetReplayMetaParam returns a new DataStoreGetReplayMetaParam
func NewDataStoreGetReplayMetaParam() *DataStoreGetReplayMetaParam {
	return &DataStoreGetReplayMetaParam{}
}
