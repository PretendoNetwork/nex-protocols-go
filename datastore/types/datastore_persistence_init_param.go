package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePersistenceInitParam is sent in the PreparePostObject method
type DataStorePersistenceInitParam struct {
	nex.Structure
	PersistenceSlotId uint16
	DeleteLastObject  bool
}

// ExtractFromStream extracts a DataStorePersistenceInitParam structure from a stream
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceInitParam.PersistenceSlotId, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.PersistenceSlotId. %s", err.Error())
	}

	dataStorePersistenceInitParam.DeleteLastObject, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.DeleteLastObject. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceInitParam
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInitParam()

	copied.PersistenceSlotId = dataStorePersistenceInitParam.PersistenceSlotId
	copied.DeleteLastObject = dataStorePersistenceInitParam.DeleteLastObject

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInitParam)

	if dataStorePersistenceInitParam.PersistenceSlotId != other.PersistenceSlotId {
		return false
	}

	if dataStorePersistenceInitParam.DeleteLastObject != other.DeleteLastObject {
		return false
	}

	return true
}

// NewDataStorePersistenceInitParam returns a new DataStorePersistenceInitParam
func NewDataStorePersistenceInitParam() *DataStorePersistenceInitParam {
	return &DataStorePersistenceInitParam{}
}
