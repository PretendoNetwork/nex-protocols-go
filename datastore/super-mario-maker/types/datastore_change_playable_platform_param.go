// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreChangePlayablePlatformParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreChangePlayablePlatformParam struct {
	nex.Structure
	DataID           uint64
	PlayablePlatform uint32
}

// ExtractFromStream extracts a DataStoreChangePlayablePlatformParam structure from a stream
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreChangePlayablePlatformParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.DataID from stream. %s", err.Error())
	}

	dataStoreChangePlayablePlatformParam.PlayablePlatform, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.PlayablePlatform from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreChangePlayablePlatformParam and returns a byte array
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreChangePlayablePlatformParam.DataID)
	stream.WriteUInt32LE(dataStoreChangePlayablePlatformParam.PlayablePlatform)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreChangePlayablePlatformParam
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangePlayablePlatformParam()

	copied.DataID = dataStoreChangePlayablePlatformParam.DataID
	copied.PlayablePlatform = dataStoreChangePlayablePlatformParam.PlayablePlatform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangePlayablePlatformParam)

	if dataStoreChangePlayablePlatformParam.DataID != other.DataID {
		return false
	}

	if dataStoreChangePlayablePlatformParam.PlayablePlatform != other.PlayablePlatform {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) String() string {
	return dataStoreChangePlayablePlatformParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangePlayablePlatformParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreChangePlayablePlatformParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreChangePlayablePlatformParam.DataID))
	b.WriteString(fmt.Sprintf("%sPlayablePlatform: %d,\n", indentationValues, dataStoreChangePlayablePlatformParam.PlayablePlatform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangePlayablePlatformParam returns a new DataStoreChangePlayablePlatformParam
func NewDataStoreChangePlayablePlatformParam() *DataStoreChangePlayablePlatformParam {
	return &DataStoreChangePlayablePlatformParam{}
}
