package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreGetSpecificMetaParam struct {
	nex.Structure
	DataIDs []uint64
}

// ExtractFromStream extracts a DataStoreGetSpecificMetaParam structure from a stream
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetSpecificMetaParam.DataIDs, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParam.DataIDs. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetSpecificMetaParam and returns a byte array
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt64LE(dataStoreGetSpecificMetaParam.DataIDs)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParam
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetSpecificMetaParam()

	copied.DataIDs = make([]uint64, len(dataStoreGetSpecificMetaParam.DataIDs))

	copy(copied.DataIDs, dataStoreGetSpecificMetaParam.DataIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetSpecificMetaParam)

	if len(dataStoreGetSpecificMetaParam.DataIDs) != len(other.DataIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetSpecificMetaParam.DataIDs); i++ {
		if dataStoreGetSpecificMetaParam.DataIDs[i] != other.DataIDs[i] {
			return false
		}
	}

	return true
}

// NewDataStoreGetSpecificMetaParam returns a new DataStoreGetSpecificMetaParam
func NewDataStoreGetSpecificMetaParam() *DataStoreGetSpecificMetaParam {
	return &DataStoreGetSpecificMetaParam{}
}
