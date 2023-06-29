package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreSpecificMetaInfoV1 struct {
	nex.Structure
	DataID   uint32
	OwnerID  uint32
	Size     uint32
	DataType uint16
	Version  uint16
}

// ExtractFromStream extracts a DataStoreSpecificMetaInfoV1 structure from a stream
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSpecificMetaInfoV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.DataID. %s", err.Error())
	}

	dataStoreSpecificMetaInfoV1.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.OwnerID. %s", err.Error())
	}

	dataStoreSpecificMetaInfoV1.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.Size. %s", err.Error())
	}

	dataStoreSpecificMetaInfoV1.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.DataType. %s", err.Error())
	}

	dataStoreSpecificMetaInfoV1.Version, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.Version. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreSpecificMetaInfoV1 and returns a byte array
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.DataID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.OwnerID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.Size)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfoV1.DataType)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfoV1.Version)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfoV1
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreSpecificMetaInfoV1()

	copied.DataID = dataStoreSpecificMetaInfoV1.DataID
	copied.OwnerID = dataStoreSpecificMetaInfoV1.OwnerID
	copied.Size = dataStoreSpecificMetaInfoV1.Size
	copied.DataType = dataStoreSpecificMetaInfoV1.DataType
	copied.Version = dataStoreSpecificMetaInfoV1.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSpecificMetaInfoV1)

	if dataStoreSpecificMetaInfoV1.DataID != other.DataID {
		return false
	}

	if dataStoreSpecificMetaInfoV1.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSpecificMetaInfoV1.Size != other.Size {
		return false
	}

	if dataStoreSpecificMetaInfoV1.DataType != other.DataType {
		return false
	}

	if dataStoreSpecificMetaInfoV1.Version != other.Version {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) String() string {
	return dataStoreSpecificMetaInfoV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSpecificMetaInfoV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreSpecificMetaInfoV1.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreSpecificMetaInfoV1.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreSpecificMetaInfoV1.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSpecificMetaInfoV1.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d\n", indentationValues, dataStoreSpecificMetaInfoV1.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfoV1 returns a new DataStoreSpecificMetaInfoV1
func NewDataStoreSpecificMetaInfoV1() *DataStoreSpecificMetaInfoV1 {
	return &DataStoreSpecificMetaInfoV1{}
}
