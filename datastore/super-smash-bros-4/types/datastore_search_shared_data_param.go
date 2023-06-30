// Package datastore_super_smash_bros_4_types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSearchSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreSearchSharedDataParam struct {
	nex.Structure
	DataType    uint8
	Owner       uint32
	Region      uint8
	Attribute1  uint8
	Attribute2  uint8
	Fighter     uint8
	ResultRange *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreSearchSharedDataParam structure from a stream
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSearchSharedDataParam.DataType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.DataType. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.Owner, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Owner. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.Region, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Region. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.Attribute1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute1. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.Attribute2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute2. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.Fighter, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Fighter. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.ResultRange. %s", err.Error())
	}

	dataStoreSearchSharedDataParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreSearchSharedDataParam and returns a byte array
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreSearchSharedDataParam.DataType)
	stream.WriteUInt32LE(dataStoreSearchSharedDataParam.Owner)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Region)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Attribute1)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Attribute2)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Fighter)
	stream.WriteStructure(dataStoreSearchSharedDataParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchSharedDataParam
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchSharedDataParam()

	copied.DataType = dataStoreSearchSharedDataParam.DataType
	copied.Owner = dataStoreSearchSharedDataParam.Owner
	copied.Region = dataStoreSearchSharedDataParam.Region
	copied.Attribute1 = dataStoreSearchSharedDataParam.Attribute1
	copied.Attribute2 = dataStoreSearchSharedDataParam.Attribute2
	copied.Fighter = dataStoreSearchSharedDataParam.Fighter
	copied.ResultRange = dataStoreSearchSharedDataParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchSharedDataParam)

	if dataStoreSearchSharedDataParam.DataType != other.DataType {
		return false
	}

	if dataStoreSearchSharedDataParam.Owner != other.Owner {
		return false
	}

	if dataStoreSearchSharedDataParam.Region != other.Region {
		return false
	}

	if dataStoreSearchSharedDataParam.Attribute1 != other.Attribute1 {
		return false
	}

	if dataStoreSearchSharedDataParam.Attribute2 != other.Attribute2 {
		return false
	}

	if dataStoreSearchSharedDataParam.Fighter != other.Fighter {
		return false
	}

	if !dataStoreSearchSharedDataParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) String() string {
	return dataStoreSearchSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSearchSharedDataParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSearchSharedDataParam.DataType))
	b.WriteString(fmt.Sprintf("%sOwner: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Owner))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %d,\n", indentationValues, dataStoreSearchSharedDataParam.Fighter))

	if dataStoreSearchSharedDataParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreSearchSharedDataParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchSharedDataParam returns a new DataStoreSearchSharedDataParam
func NewDataStoreSearchSharedDataParam() *DataStoreSearchSharedDataParam {
	return &DataStoreSearchSharedDataParam{}
}
