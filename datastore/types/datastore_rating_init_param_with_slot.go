// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

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

// String returns a string representation of the struct
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) String() string {
	return dataStoreRatingInitParamWithSlot.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInitParamWithSlot{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRatingInitParamWithSlot.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreRatingInitParamWithSlot.Slot))

	if dataStoreRatingInitParamWithSlot.Param != nil {
		b.WriteString(fmt.Sprintf("%sParam: %s\n", indentationValues, dataStoreRatingInitParamWithSlot.Param.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sParam: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	return &DataStoreRatingInitParamWithSlot{}
}
