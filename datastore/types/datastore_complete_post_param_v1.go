// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreCompletePostParamV1 is a data structure used by the DataStore protocol
type DataStoreCompletePostParamV1 struct {
	nex.Structure
	DataID    uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompletePostParamV1 structure from a stream
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostParamV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.DataID. %s", err.Error())
	}

	dataStoreCompletePostParamV1.IsSuccess, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.IsSuccess. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCompletePostParamV1 and returns a byte array
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCompletePostParamV1.DataID)
	stream.WriteBool(dataStoreCompletePostParamV1.IsSuccess)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostParamV1
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostParamV1()

	copied.SetStructureVersion(dataStoreCompletePostParamV1.StructureVersion())

	copied.DataID = dataStoreCompletePostParamV1.DataID
	copied.IsSuccess = dataStoreCompletePostParamV1.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParamV1)

	if dataStoreCompletePostParamV1.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreCompletePostParamV1.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParamV1.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) String() string {
	return dataStoreCompletePostParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompletePostParamV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreCompletePostParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %t\n", indentationValues, dataStoreCompletePostParamV1.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParamV1 returns a new DataStoreCompletePostParamV1
func NewDataStoreCompletePostParamV1() *DataStoreCompletePostParamV1 {
	return &DataStoreCompletePostParamV1{
		DataID:    0,
		IsSuccess: false,
	}
}
