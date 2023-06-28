package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreCompleteUpdateParam struct {
	nex.Structure
	DataID    uint64
	Version   uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompleteUpdateParam structure from a stream
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompleteUpdateParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
	}

	dataStoreCompleteUpdateParam.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.Version. %s", err.Error())
	}

	dataStoreCompleteUpdateParam.IsSuccess, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCompleteUpdateParam and returns a byte array
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompleteUpdateParam.DataID)
	stream.WriteUInt32LE(dataStoreCompleteUpdateParam.Version)
	stream.WriteBool(dataStoreCompleteUpdateParam.IsSuccess)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompleteUpdateParam
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompleteUpdateParam()

	copied.DataID = dataStoreCompleteUpdateParam.DataID
	copied.Version = dataStoreCompleteUpdateParam.Version
	copied.IsSuccess = dataStoreCompleteUpdateParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompleteUpdateParam)

	if dataStoreCompleteUpdateParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompleteUpdateParam.Version != other.Version {
		return false
	}

	if dataStoreCompleteUpdateParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	return &DataStoreCompleteUpdateParam{}
}
