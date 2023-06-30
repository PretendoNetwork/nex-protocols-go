// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreCompleteUpdateParam is a data structure used by the DataStore protocol
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

// String returns a string representation of the struct
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) String() string {
	return dataStoreCompleteUpdateParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompleteUpdateParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompleteUpdateParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreCompleteUpdateParam.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, dataStoreCompleteUpdateParam.Version))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %t\n", indentationValues, dataStoreCompleteUpdateParam.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	return &DataStoreCompleteUpdateParam{}
}
