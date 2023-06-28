package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	nex.Structure
	OwnerID           uint32
	PersistenceSlotID uint16
}

// ExtractFromStream extracts a DataStorePersistenceTarget structure from a stream
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceTarget.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.OwnerID. %s", err.Error())
	}

	dataStorePersistenceTarget.PersistenceSlotID, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.PersistenceSlotID. %s", err.Error())
	}

	return nil
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{}
}
