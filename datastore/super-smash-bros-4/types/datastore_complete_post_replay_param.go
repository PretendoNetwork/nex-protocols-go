// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCompletePostReplayParam is a data structure used by the DataStore Super Mario Maker protocol
type DataStoreCompletePostReplayParam struct {
	nex.Structure
	ReplayID      uint64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostReplayParam
}

// ExtractFromStream extracts a DataStoreCompletePostReplayParam structure from a stream
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostReplayParam.ReplayID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.ReplayID. %s", err.Error())
	}

	completeParam, err := stream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.CompleteParam. %s", err.Error())
	}

	dataStoreCompletePostReplayParam.CompleteParam = completeParam.(*datastore_types.DataStoreCompletePostParam)

	prepareParam, err := stream.ReadStructure(NewDataStorePreparePostReplayParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.PrepareParam. %s", err.Error())
	}

	dataStoreCompletePostReplayParam.PrepareParam = prepareParam.(*DataStorePreparePostReplayParam)

	return nil
}

// Bytes encodes the DataStoreCompletePostReplayParam and returns a byte array
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompletePostReplayParam.ReplayID)
	stream.WriteStructure(dataStoreCompletePostReplayParam.CompleteParam)
	stream.WriteStructure(dataStoreCompletePostReplayParam.PrepareParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostReplayParam
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostReplayParam()

	copied.ReplayID = dataStoreCompletePostReplayParam.ReplayID
	copied.CompleteParam = dataStoreCompletePostReplayParam.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostReplayParam.PrepareParam.Copy().(*DataStorePreparePostReplayParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostReplayParam)

	if dataStoreCompletePostReplayParam.ReplayID != other.ReplayID {
		return false
	}

	if !dataStoreCompletePostReplayParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostReplayParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) String() string {
	return dataStoreCompletePostReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompletePostReplayParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReplayID: %d,\n", indentationValues, dataStoreCompletePostReplayParam.ReplayID))

	if dataStoreCompletePostReplayParam.CompleteParam != nil {
		b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dataStoreCompletePostReplayParam.CompleteParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCompleteParam: nil,\n", indentationValues))
	}

	if dataStoreCompletePostReplayParam.PrepareParam != nil {
		b.WriteString(fmt.Sprintf("%sPrepareParam: %s\n", indentationValues, dataStoreCompletePostReplayParam.PrepareParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareParam: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostReplayParam returns a new DataStoreCompletePostReplayParam
func NewDataStoreCompletePostReplayParam() *DataStoreCompletePostReplayParam {
	return &DataStoreCompletePostReplayParam{}
}
