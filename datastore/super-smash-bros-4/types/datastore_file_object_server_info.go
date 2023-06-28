package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	nex.Structure
	DataID  uint64
	GetInfo *datastore_types.DataStoreReqGetInfo
}

// ExtractFromStream extracts a DataStoreFileServerObjectInfo structure from a stream
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreFileServerObjectInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.DataID. %s", err.Error())
	}

	getInfo, err := stream.ReadStructure(datastore_types.NewDataStoreReqGetInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.GetInfo. %s", err.Error())
	}

	dataStoreFileServerObjectInfo.GetInfo = getInfo.(*datastore_types.DataStoreReqGetInfo)

	return nil
}

// Bytes encodes the DataStoreFileServerObjectInfo and returns a byte array
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreFileServerObjectInfo.DataID)
	stream.WriteStructure(dataStoreFileServerObjectInfo.GetInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFileServerObjectInfo
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreFileServerObjectInfo()

	copied.DataID = dataStoreFileServerObjectInfo.DataID
	copied.GetInfo = dataStoreFileServerObjectInfo.GetInfo.Copy().(*datastore_types.DataStoreReqGetInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFileServerObjectInfo)

	if dataStoreFileServerObjectInfo.DataID != other.DataID {
		return false
	}

	if !dataStoreFileServerObjectInfo.GetInfo.Equals(other.GetInfo) {
		return false
	}

	return true
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{}
}
