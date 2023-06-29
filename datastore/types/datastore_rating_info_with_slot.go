package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingInfoWithSlot is a data structure used by the DataStore protocol
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

// String returns a string representation of the struct
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) String() string {
	return dataStoreRatingInfoWithSlot.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInfoWithSlot{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRatingInfoWithSlot.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreRatingInfoWithSlot.Slot))

	if dataStoreRatingInfoWithSlot.Rating != nil {
		b.WriteString(fmt.Sprintf("%sRating: %s\n", indentationValues, dataStoreRatingInfoWithSlot.Rating.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sRating: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	return &DataStoreRatingInfoWithSlot{}
}
