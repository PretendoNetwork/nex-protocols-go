// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

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

// String returns a string representation of the struct
func (dataStoreGetMetaParam *DataStoreGetMetaParam) String() string {
	return dataStoreGetMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetMetaParam *DataStoreGetMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetMetaParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreGetMetaParam.DataID))

	if dataStoreGetMetaParam.PersistenceTarget != nil {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dataStoreGetMetaParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sResultOption: %d,\n", indentationValues, dataStoreGetMetaParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %d\n", indentationValues, dataStoreGetMetaParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{}
}
