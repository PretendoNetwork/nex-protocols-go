// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSpecificMetaInfo is a data structure used by the DataStore protocol
type DataStoreSpecificMetaInfo struct {
	nex.Structure
	DataID   uint64
	OwnerID  uint32
	Size     uint32
	DataType uint16
	Version  uint32
}

// ExtractFromStream extracts a DataStoreSpecificMetaInfo structure from a stream
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSpecificMetaInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataID. %s", err.Error())
	}

	dataStoreSpecificMetaInfo.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.OwnerID. %s", err.Error())
	}

	dataStoreSpecificMetaInfo.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Size. %s", err.Error())
	}

	dataStoreSpecificMetaInfo.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataType. %s", err.Error())
	}

	dataStoreSpecificMetaInfo.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Version. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreSpecificMetaInfo and returns a byte array
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreSpecificMetaInfo.DataID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.OwnerID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.Size)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfo.DataType)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.Version)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfo
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreSpecificMetaInfo()

	copied.DataID = dataStoreSpecificMetaInfo.DataID
	copied.OwnerID = dataStoreSpecificMetaInfo.OwnerID
	copied.Size = dataStoreSpecificMetaInfo.Size
	copied.DataType = dataStoreSpecificMetaInfo.DataType
	copied.Version = dataStoreSpecificMetaInfo.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSpecificMetaInfo)

	if dataStoreSpecificMetaInfo.DataID != other.DataID {
		return false
	}

	if dataStoreSpecificMetaInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSpecificMetaInfo.Size != other.Size {
		return false
	}

	if dataStoreSpecificMetaInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSpecificMetaInfo.Version != other.Version {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) String() string {
	return dataStoreSpecificMetaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSpecificMetaInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreSpecificMetaInfo.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreSpecificMetaInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreSpecificMetaInfo.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSpecificMetaInfo.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d\n", indentationValues, dataStoreSpecificMetaInfo.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfo returns a new DataStoreSpecificMetaInfo
func NewDataStoreSpecificMetaInfo() *DataStoreSpecificMetaInfo {
	return &DataStoreSpecificMetaInfo{}
}
