package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

type DataStoreCompletePostSharedDataParam struct {
	nex.Structure
	DataID        uint64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// ExtractFromStream extracts a DataStoreCompletePostSharedDataParam structure from a stream
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostSharedDataParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.DataID. %s", err.Error())
	}

	completeParam, err := stream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.CompleteParam. %s", err.Error())
	}

	dataStoreCompletePostSharedDataParam.CompleteParam = completeParam.(*datastore_types.DataStoreCompletePostParam)

	prepareParam, err := stream.ReadStructure(NewDataStorePreparePostSharedDataParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.PrepareParam. %s", err.Error())
	}

	dataStoreCompletePostSharedDataParam.PrepareParam = prepareParam.(*DataStorePreparePostSharedDataParam)

	return nil
}

// Bytes encodes the DataStoreCompletePostSharedDataParam and returns a byte array
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompletePostSharedDataParam.DataID)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.CompleteParam)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.PrepareParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostSharedDataParam
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostSharedDataParam()

	copied.DataID = dataStoreCompletePostSharedDataParam.DataID
	copied.CompleteParam = dataStoreCompletePostSharedDataParam.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostSharedDataParam.PrepareParam.Copy().(*DataStorePreparePostSharedDataParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostSharedDataParam)

	if dataStoreCompletePostSharedDataParam.DataID != other.DataID {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// NewDataStoreCompletePostSharedDataParam returns a new DataStoreCompletePostSharedDataParam
func NewDataStoreCompletePostSharedDataParam() *DataStoreCompletePostSharedDataParam {
	return &DataStoreCompletePostSharedDataParam{}
}
