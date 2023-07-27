// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePrepareGetParam is sent in the PrepareGetObject method
type DataStorePrepareGetParam struct {
	nex.Structure
	DataID            uint64
	LockID            uint32
	PersistenceTarget *DataStorePersistenceTarget
	AccessPassword    uint64
	ExtraData         []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePrepareGetParam structure from a stream
func (dataStorePrepareGetParam *DataStorePrepareGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	dataStorePrepareGetParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.DataID. %s", err.Error())
	}

	dataStorePrepareGetParam.LockID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.LockID. %s", err.Error())
	}

	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.PersistenceTarget. %s", err.Error())
	}

	dataStorePrepareGetParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStorePrepareGetParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.AccessPassword. %s", err.Error())
	}

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		dataStorePrepareGetParam.ExtraData, err = stream.ReadListString()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareGetParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetParam
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareGetParam()

	copied.DataID = dataStorePrepareGetParam.DataID
	copied.LockID = dataStorePrepareGetParam.LockID
	copied.PersistenceTarget = dataStorePrepareGetParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.AccessPassword = dataStorePrepareGetParam.AccessPassword
	copied.ExtraData = make([]string, len(dataStorePrepareGetParam.ExtraData))

	copy(copied.ExtraData, dataStorePrepareGetParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetParam)

	if dataStorePrepareGetParam.DataID != other.DataID {
		return false
	}

	if dataStorePrepareGetParam.LockID != other.LockID {
		return false
	}

	if !dataStorePrepareGetParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if dataStorePrepareGetParam.AccessPassword != other.AccessPassword {
		return false
	}

	if len(dataStorePrepareGetParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePrepareGetParam.ExtraData); i++ {
		if dataStorePrepareGetParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePrepareGetParam *DataStorePrepareGetParam) String() string {
	return dataStorePrepareGetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareGetParam *DataStorePrepareGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePrepareGetParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStorePrepareGetParam.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %d,\n", indentationValues, dataStorePrepareGetParam.LockID))

	if dataStorePrepareGetParam.PersistenceTarget != nil {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dataStorePrepareGetParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sAccessPassword: %d,\n", indentationValues, dataStorePrepareGetParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %v\n", indentationValues, dataStorePrepareGetParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParam returns a new DataStorePrepareGetParam
func NewDataStorePrepareGetParam() *DataStorePrepareGetParam {
	return &DataStorePrepareGetParam{}
}
