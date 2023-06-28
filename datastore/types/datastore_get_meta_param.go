package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetMetaParam is sent in the GetMeta method
type DataStoreGetMetaParam struct {
	nex.Structure
	DataID            uint64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      uint8
	AccessPassword    uint64
}

// ExtractFromStream extracts a DataStoreGetMetaParam structure from a stream
func (dataStoreGetMetaParam *DataStoreGetMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetMetaParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.DataID. %s", err.Error())
	}

	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.PersistenceTarget. %s", err.Error())
	}

	dataStoreGetMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStoreGetMetaParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.ResultOption. %s", err.Error())
	}

	dataStoreGetMetaParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetMetaParam
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetMetaParam()

	copied.DataID = dataStoreGetMetaParam.DataID
	copied.PersistenceTarget = dataStoreGetMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.ResultOption = dataStoreGetMetaParam.ResultOption
	copied.AccessPassword = dataStoreGetMetaParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetMetaParam)

	if dataStoreGetMetaParam.DataID != other.DataID {
		return false
	}

	if !dataStoreGetMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if dataStoreGetMetaParam.ResultOption != other.ResultOption {
		return false
	}

	if dataStoreGetMetaParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{}
}
