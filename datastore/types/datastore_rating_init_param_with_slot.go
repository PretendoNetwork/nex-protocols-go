package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingInitParamWithSlot is sent in the PreparePostObject method
type DataStoreRatingInitParamWithSlot struct {
	nex.Structure
	Slot  int8
	Param *DataStoreRatingInitParam
}

// ExtractFromStream extracts a DataStoreRatingInitParamWithSlot structure from a stream
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingInitParamWithSlot.Slot, err = stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Slot. %s", err.Error())
	}

	param, err := stream.ReadStructure(NewDataStoreRatingInitParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Param. %s", err.Error())
	}

	dataStoreRatingInitParamWithSlot.Param = param.(*DataStoreRatingInitParam)

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParamWithSlot
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInitParamWithSlot()

	copied.Slot = dataStoreRatingInitParamWithSlot.Slot
	copied.Param = dataStoreRatingInitParamWithSlot.Param.Copy().(*DataStoreRatingInitParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInitParamWithSlot)

	if dataStoreRatingInitParamWithSlot.Slot != other.Slot {
		return false
	}

	if !dataStoreRatingInitParamWithSlot.Param.Equals(other.Param) {
		return false
	}

	return true
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	return &DataStoreRatingInitParamWithSlot{}
}
