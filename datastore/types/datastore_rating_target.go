package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingTarget is sent in the RateObjects method
type DataStoreRatingTarget struct {
	nex.Structure
	DataID uint64
	Slot   uint8
}

// ExtractFromStream extracts a DataStoreRatingTarget structure from a stream
func (dataStoreRatingTarget *DataStoreRatingTarget) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingTarget.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.DataID. %s", err.Error())
	}

	dataStoreRatingTarget.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingTarget
func (dataStoreRatingTarget *DataStoreRatingTarget) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingTarget()

	copied.DataID = dataStoreRatingTarget.DataID
	copied.Slot = dataStoreRatingTarget.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingTarget *DataStoreRatingTarget) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingTarget)

	if dataStoreRatingTarget.DataID != other.DataID {
		return false
	}

	if dataStoreRatingTarget.Slot != other.Slot {
		return false
	}

	return true
}

// NewDataStoreRatingTarget returns a new DataStoreRatingTarget
func NewDataStoreRatingTarget() *DataStoreRatingTarget {
	return &DataStoreRatingTarget{}
}
