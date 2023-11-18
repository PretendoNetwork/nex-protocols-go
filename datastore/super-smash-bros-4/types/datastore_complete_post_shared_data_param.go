// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCompletePostSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreCompletePostSharedDataParam struct {
	nex.Structure
	DataID        uint64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// ExtractFromStream extracts a DataStoreCompletePostSharedDataParam structure from a stream
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostSharedDataParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.DataID. %s", err.Error())
	}

	dataStoreCompletePostSharedDataParam.CompleteParam, err = nex.StreamReadStructure(stream, datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.CompleteParam. %s", err.Error())
	}

	dataStoreCompletePostSharedDataParam.PrepareParam, err = nex.StreamReadStructure(stream, NewDataStorePreparePostSharedDataParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.PrepareParam. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCompletePostSharedDataParam and returns a byte array
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompletePostSharedDataParam.DataID)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.CompleteParam)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.PrepareParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostSharedDataParam
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostSharedDataParam()

	copied.SetStructureVersion(dataStoreCompletePostSharedDataParam.StructureVersion())

	copied.DataID = dataStoreCompletePostSharedDataParam.DataID
	copied.CompleteParam = dataStoreCompletePostSharedDataParam.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostSharedDataParam.PrepareParam.Copy().(*DataStorePreparePostSharedDataParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostSharedDataParam)

	if dataStoreCompletePostSharedDataParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreCompletePostSharedDataParam.DataID != other.DataID {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) String() string {
	return dataStoreCompletePostSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompletePostSharedDataParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreCompletePostSharedDataParam.DataID))

	if dataStoreCompletePostSharedDataParam.CompleteParam != nil {
		b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dataStoreCompletePostSharedDataParam.CompleteParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCompleteParam: nil,\n", indentationValues))
	}

	if dataStoreCompletePostSharedDataParam.PrepareParam != nil {
		b.WriteString(fmt.Sprintf("%sPrepareParam: %s\n", indentationValues, dataStoreCompletePostSharedDataParam.PrepareParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareParam: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostSharedDataParam returns a new DataStoreCompletePostSharedDataParam
func NewDataStoreCompletePostSharedDataParam() *DataStoreCompletePostSharedDataParam {
	return &DataStoreCompletePostSharedDataParam{}
}
