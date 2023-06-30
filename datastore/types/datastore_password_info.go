// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePasswordInfo is a data structure used by the DataStore protocol
type DataStorePasswordInfo struct {
	nex.Structure
	DataID         uint64
	AccessPassword uint64
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStorePasswordInfo structure from a stream
func (dataStorePasswordInfo *DataStorePasswordInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePasswordInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.DataID. %s", err.Error())
	}

	dataStorePasswordInfo.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.AccessPassword. %s", err.Error())
	}

	dataStorePasswordInfo.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePasswordInfo and returns a byte array
func (dataStorePasswordInfo *DataStorePasswordInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStorePasswordInfo.DataID)
	stream.WriteUInt64LE(dataStorePasswordInfo.AccessPassword)
	stream.WriteUInt64LE(dataStorePasswordInfo.UpdatePassword)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePasswordInfo
func (dataStorePasswordInfo *DataStorePasswordInfo) Copy() nex.StructureInterface {
	copied := NewDataStorePasswordInfo()

	copied.DataID = dataStorePasswordInfo.DataID
	copied.AccessPassword = dataStorePasswordInfo.AccessPassword
	copied.UpdatePassword = dataStorePasswordInfo.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePasswordInfo *DataStorePasswordInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePasswordInfo)

	if dataStorePasswordInfo.DataID != other.DataID {
		return false
	}

	if dataStorePasswordInfo.AccessPassword != other.AccessPassword {
		return false
	}

	if dataStorePasswordInfo.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePasswordInfo *DataStorePasswordInfo) String() string {
	return dataStorePasswordInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePasswordInfo *DataStorePasswordInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePasswordInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePasswordInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStorePasswordInfo.DataID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %d,\n", indentationValues, dataStorePasswordInfo.AccessPassword))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %d\n", indentationValues, dataStorePasswordInfo.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePasswordInfo returns a new DataStorePasswordInfo
func NewDataStorePasswordInfo() *DataStorePasswordInfo {
	return &DataStorePasswordInfo{}
}
