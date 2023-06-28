package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreRatingInfoWithSlot struct {
	nex.Structure
	Slot   int8
	Rating *DataStoreRatingInfo
}

// ExtractFromStream extracts a DataStoreRatingInfoWithSlot structure from a stream
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingInfoWithSlot.Slot, err = stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Slot. %s", err.Error())
	}

	rating, err := stream.ReadStructure(NewDataStoreRatingInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Rating. %s", err.Error())
	}

	dataStoreRatingInfoWithSlot.Rating = rating.(*DataStoreRatingInfo)

	return nil
}

// Bytes encodes the DataStoreRatingInfoWithSlot and returns a byte array
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(uint8(dataStoreRatingInfoWithSlot.Slot))
	stream.WriteStructure(dataStoreRatingInfoWithSlot.Rating)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRatingInfoWithSlot
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInfoWithSlot()

	copied.Slot = dataStoreRatingInfoWithSlot.Slot
	copied.Rating = dataStoreRatingInfoWithSlot.Rating.Copy().(*DataStoreRatingInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInfoWithSlot)

	if dataStoreRatingInfoWithSlot.Slot != other.Slot {
		return false
	}

	if !dataStoreRatingInfoWithSlot.Rating.Equals(other.Rating) {
		return false
	}

	return true
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	return &DataStoreRatingInfoWithSlot{}
}
