// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePreparePostSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePreparePostSharedDataParam struct {
	nex.Structure
	DataType   uint8
	Region     uint8
	Attribute1 uint8
	Attribute2 uint8
	Fighter    []byte
	Size       uint32
	Comment    string
	MetaBinary []byte
	ExtraData  []string
}

// ExtractFromStream extracts a DataStorePreparePostSharedDataParam structure from a stream
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePreparePostSharedDataParam.DataType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.DataType. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Region, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Region. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Attribute1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute1. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Attribute2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute2. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Fighter, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Fighter. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Size. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.Comment, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Comment. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.MetaBinary. %s", err.Error())
	}

	dataStorePreparePostSharedDataParam.ExtraData, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePreparePostSharedDataParam and returns a byte array
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.DataType)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Region)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Attribute1)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Attribute2)
	stream.WriteBuffer(dataStorePreparePostSharedDataParam.Fighter)
	stream.WriteUInt32LE(dataStorePreparePostSharedDataParam.Size)
	stream.WriteString(dataStorePreparePostSharedDataParam.Comment)
	stream.WriteQBuffer(dataStorePreparePostSharedDataParam.MetaBinary)
	stream.WriteListString(dataStorePreparePostSharedDataParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePreparePostSharedDataParam
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostSharedDataParam()

	copied.SetStructureVersion(dataStorePreparePostSharedDataParam.StructureVersion())

	copied.DataType = dataStorePreparePostSharedDataParam.DataType
	copied.Region = dataStorePreparePostSharedDataParam.Region
	copied.Attribute1 = dataStorePreparePostSharedDataParam.Attribute1
	copied.Attribute2 = dataStorePreparePostSharedDataParam.Attribute2
	copied.Fighter = make([]byte, len(dataStorePreparePostSharedDataParam.Fighter))

	copy(copied.Fighter, dataStorePreparePostSharedDataParam.Fighter)

	copied.Size = dataStorePreparePostSharedDataParam.Size
	copied.Comment = dataStorePreparePostSharedDataParam.Comment
	copied.MetaBinary = make([]byte, len(dataStorePreparePostSharedDataParam.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostSharedDataParam.MetaBinary)

	copied.ExtraData = make([]string, len(dataStorePreparePostSharedDataParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostSharedDataParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostSharedDataParam)

	if dataStorePreparePostSharedDataParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStorePreparePostSharedDataParam.DataType != other.DataType {
		return false
	}

	if dataStorePreparePostSharedDataParam.Region != other.Region {
		return false
	}

	if dataStorePreparePostSharedDataParam.Attribute1 != other.Attribute1 {
		return false
	}

	if dataStorePreparePostSharedDataParam.Attribute2 != other.Attribute2 {
		return false
	}

	if !bytes.Equal(dataStorePreparePostSharedDataParam.Fighter, other.Fighter) {
		return false
	}

	if dataStorePreparePostSharedDataParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostSharedDataParam.Comment != other.Comment {
		return false
	}

	if !bytes.Equal(dataStorePreparePostSharedDataParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStorePreparePostSharedDataParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostSharedDataParam.ExtraData); i++ {
		if dataStorePreparePostSharedDataParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) String() string {
	return dataStorePreparePostSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.DataType))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %x,\n", indentationValues, dataStorePreparePostSharedDataParam.Fighter))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStorePreparePostSharedDataParam.Size))
	b.WriteString(fmt.Sprintf("%sComment: %q,\n", indentationValues, dataStorePreparePostSharedDataParam.Comment))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStorePreparePostSharedDataParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sExtraData: %v\n", indentationValues, dataStorePreparePostSharedDataParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostSharedDataParam returns a new DataStorePreparePostSharedDataParam
func NewDataStorePreparePostSharedDataParam() *DataStorePreparePostSharedDataParam {
	return &DataStorePreparePostSharedDataParam{}
}
