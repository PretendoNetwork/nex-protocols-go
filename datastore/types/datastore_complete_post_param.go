package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreCompletePostParam is sent in the CompletePostObject method
type DataStoreCompletePostParam struct {
	nex.Structure
	DataID    uint64
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompletePostParam structure from a stream
func (dataStoreCompletePostParam *DataStoreCompletePostParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.DataID. %s", err.Error())
	}

	dataStoreCompletePostParam.IsSuccess, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostParam
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostParam()

	copied.DataID = dataStoreCompletePostParam.DataID
	copied.IsSuccess = dataStoreCompletePostParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParam)

	if dataStoreCompletePostParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// NewDataStoreCompletePostParam returns a new DataStoreCompletePostParam
func NewDataStoreCompletePostParam() *DataStoreCompletePostParam {
	return &DataStoreCompletePostParam{}
}
