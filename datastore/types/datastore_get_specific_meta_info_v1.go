package datastore_types

import (
	"fmt"

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

// NewDataStoreSpecificMetaInfoV1 returns a new DataStoreSpecificMetaInfoV1
func NewDataStoreSpecificMetaInfoV1() *DataStoreSpecificMetaInfoV1 {
	return &DataStoreSpecificMetaInfoV1{}
}
