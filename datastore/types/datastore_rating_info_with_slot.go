// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingInfoWithSlot is a data structure used by the DataStore protocol
type DataStoreRatingInfoWithSlot struct {
	types.Structure
	Slot   *types.PrimitiveS8
	Rating *DataStoreRatingInfo
}

// ExtractFrom extracts the DataStoreRatingInfoWithSlot from the given readable
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingInfoWithSlot.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingInfoWithSlot header. %s", err.Error())
	}

	err = dataStoreRatingInfoWithSlot.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Slot. %s", err.Error())
	}

	err = dataStoreRatingInfoWithSlot.Rating.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Rating. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreRatingInfoWithSlot to the given writable
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingInfoWithSlot.Slot.WriteTo(contentWritable)
	dataStoreRatingInfoWithSlot.Rating.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingInfoWithSlot.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreRatingInfoWithSlot
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Copy() types.RVType {
	copied := NewDataStoreRatingInfoWithSlot()

	copied.StructureVersion = dataStoreRatingInfoWithSlot.StructureVersion

	copied.Slot = dataStoreRatingInfoWithSlot.Slot.Copy().(*types.PrimitiveS8)
	copied.Rating = dataStoreRatingInfoWithSlot.Rating.Copy().(*DataStoreRatingInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInfoWithSlot); !ok {
		return false
	}

	other := o.(*DataStoreRatingInfoWithSlot)

	if dataStoreRatingInfoWithSlot.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingInfoWithSlot.Slot.Equals(other.Slot) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingInfoWithSlot.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreRatingInfoWithSlot.Slot))
	b.WriteString(fmt.Sprintf("%sRating: %s\n", indentationValues, dataStoreRatingInfoWithSlot.Rating.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	return &DataStoreRatingInfoWithSlot{
		Slot:   types.NewPrimitiveS8(0),
		Rating: NewDataStoreRatingInfo(),
	}
}
