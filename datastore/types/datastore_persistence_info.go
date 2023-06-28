package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStorePersistenceInfo struct {
	nex.Structure
	OwnerID           uint32
	PersistenceSlotID uint16
	DataID            uint64
}

// ExtractFromStream extracts a DataStorePersistenceInfo structure from a stream
func (dataStorePersistenceInfo *DataStorePersistenceInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceInfo.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.OwnerID. %s", err.Error())
	}

	dataStorePersistenceInfo.PersistenceSlotID, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.PersistenceSlotID. %s", err.Error())
	}

	dataStorePersistenceInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.DataID. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePersistenceInfo and returns a byte array
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePersistenceInfo.OwnerID)
	stream.WriteUInt16LE(dataStorePersistenceInfo.PersistenceSlotID)
	stream.WriteUInt64LE(dataStorePersistenceInfo.DataID)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePersistenceInfo
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInfo()

	copied.OwnerID = dataStorePersistenceInfo.OwnerID
	copied.PersistenceSlotID = dataStorePersistenceInfo.PersistenceSlotID
	copied.DataID = dataStorePersistenceInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInfo)

	if dataStorePersistenceInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStorePersistenceInfo.PersistenceSlotID != other.PersistenceSlotID {
		return false
	}

	if dataStorePersistenceInfo.DataID != other.DataID {
		return false
	}

	return true
}

// NewDataStorePersistenceInfo returns a new DataStorePersistenceInfo
func NewDataStorePersistenceInfo() *DataStorePersistenceInfo {
	return &DataStorePersistenceInfo{}
}
