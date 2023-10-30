// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqGetAdditionalMeta is a data structure used by the DataStore protocol
type DataStoreReqGetAdditionalMeta struct {
	nex.Structure
	OwnerID    uint32
	DataType   uint16
	Version    uint16
	MetaBinary []byte
}

// ExtractFromStream extracts a DataStoreReqGetAdditionalMeta structure from a stream
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqGetAdditionalMeta.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.OwnerID. %s", err.Error())
	}

	dataStoreReqGetAdditionalMeta.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.DataType. %s", err.Error())
	}

	dataStoreReqGetAdditionalMeta.Version, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.Version. %s", err.Error())
	}

	dataStoreReqGetAdditionalMeta.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.MetaBinary. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqGetAdditionalMeta and returns a byte array
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqGetAdditionalMeta.OwnerID)
	stream.WriteUInt16LE(dataStoreReqGetAdditionalMeta.DataType)
	stream.WriteUInt16LE(dataStoreReqGetAdditionalMeta.Version)
	stream.WriteQBuffer(dataStoreReqGetAdditionalMeta.MetaBinary)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.SetStructureVersion(dataStoreReqGetAdditionalMeta.StructureVersion())

	copied.OwnerID = dataStoreReqGetAdditionalMeta.OwnerID
	copied.DataType = dataStoreReqGetAdditionalMeta.DataType
	copied.Version = dataStoreReqGetAdditionalMeta.Version
	copied.MetaBinary = make([]byte, len(dataStoreReqGetAdditionalMeta.MetaBinary))

	copy(copied.MetaBinary, dataStoreReqGetAdditionalMeta.MetaBinary)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetAdditionalMeta)

	if dataStoreReqGetAdditionalMeta.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreReqGetAdditionalMeta.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreReqGetAdditionalMeta.DataType != other.DataType {
		return false
	}

	if dataStoreReqGetAdditionalMeta.Version != other.Version {
		return false
	}

	if !bytes.Equal(dataStoreReqGetAdditionalMeta.MetaBinary, other.MetaBinary) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) String() string {
	return dataStoreReqGetAdditionalMeta.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetAdditionalMeta{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.Version))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x\n", indentationValues, dataStoreReqGetAdditionalMeta.MetaBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	return &DataStoreReqGetAdditionalMeta{
		OwnerID:    0,
		DataType:   0,
		Version:    0,
		MetaBinary: make([]byte, 0),
	}
}
