// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

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

	copied.SetStructureVersion(dataStoreCompletePostParam.StructureVersion())

	copied.DataID = dataStoreCompletePostParam.DataID
	copied.IsSuccess = dataStoreCompletePostParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParam)

	if dataStoreCompletePostParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreCompletePostParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostParam *DataStoreCompletePostParam) String() string {
	return dataStoreCompletePostParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostParam *DataStoreCompletePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompletePostParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreCompletePostParam.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %t\n", indentationValues, dataStoreCompletePostParam.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParam returns a new DataStoreCompletePostParam
func NewDataStoreCompletePostParam() *DataStoreCompletePostParam {
	return &DataStoreCompletePostParam{}
}
