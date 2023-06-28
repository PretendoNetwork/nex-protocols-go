package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

// NewDataStoreSearchSharedDataParam returns a new DataStoreSearchSharedDataParam
func NewDataStoreSearchSharedDataParam() *DataStoreSearchSharedDataParam {
	return &DataStoreSearchSharedDataParam{}
}
