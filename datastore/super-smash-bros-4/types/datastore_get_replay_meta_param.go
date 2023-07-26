// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetReplayMetaParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreGetReplayMetaParam struct {
	nex.Structure
	ReplayID uint64
	MetaType uint8
}

// ExtractFromStream extracts a DataStoreGetReplayMetaParam structure from a stream
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetReplayMetaParam.ReplayID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.ReplayID. %s", err.Error())
	}

	dataStoreGetReplayMetaParam.MetaType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.MetaType. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetReplayMetaParam and returns a byte array
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetReplayMetaParam.ReplayID)
	stream.WriteUInt8(dataStoreGetReplayMetaParam.MetaType)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetReplayMetaParam
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetReplayMetaParam()

	copied.ReplayID = dataStoreGetReplayMetaParam.ReplayID
	copied.MetaType = dataStoreGetReplayMetaParam.MetaType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetReplayMetaParam)

	if dataStoreGetReplayMetaParam.ReplayID != other.ReplayID {
		return false
	}

	if dataStoreGetReplayMetaParam.MetaType != other.MetaType {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) String() string {
	return dataStoreGetReplayMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetReplayMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetReplayMetaParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReplayID: %d,\n", indentationValues, dataStoreGetReplayMetaParam.ReplayID))
	b.WriteString(fmt.Sprintf("%sMetaType: %d\n", indentationValues, dataStoreGetReplayMetaParam.MetaType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetReplayMetaParam returns a new DataStoreGetReplayMetaParam
func NewDataStoreGetReplayMetaParam() *DataStoreGetReplayMetaParam {
	return &DataStoreGetReplayMetaParam{}
}
