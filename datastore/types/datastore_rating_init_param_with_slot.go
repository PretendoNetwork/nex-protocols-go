// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingInitParamWithSlot is sent in the PreparePostObject method
type DataStoreRatingInitParamWithSlot struct {
	types.Structure
	Slot  *types.PrimitiveS8
	Param *DataStoreRatingInitParam
}

// WriteTo writes the DataStoreRatingInitParamWithSlot to the given writable
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingInitParamWithSlot.Slot.WriteTo(contentWritable)
	dataStoreRatingInitParamWithSlot.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingInitParamWithSlot.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInitParamWithSlot from the given readable
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingInitParamWithSlot.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingInitParamWithSlot header. %s", err.Error())
	}

	err = dataStoreRatingInitParamWithSlot.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Slot. %s", err.Error())
	}

	err = dataStoreRatingInitParamWithSlot.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Param. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParamWithSlot
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Copy() types.RVType {
	copied := NewDataStoreRatingInitParamWithSlot()

	copied.StructureVersion = dataStoreRatingInitParamWithSlot.StructureVersion

	copied.Slot = dataStoreRatingInitParamWithSlot.Slot.Copy().(*types.PrimitiveS8)
	copied.Param = dataStoreRatingInitParamWithSlot.Param.Copy().(*DataStoreRatingInitParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInitParamWithSlot); !ok {
		return false
	}

	other := o.(*DataStoreRatingInitParamWithSlot)

	if dataStoreRatingInitParamWithSlot.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingInitParamWithSlot.Slot.Equals(other.Slot) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingInitParamWithSlot.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreRatingInitParamWithSlot.Slot))
	b.WriteString(fmt.Sprintf("%sParam: %s\n", indentationValues, dataStoreRatingInitParamWithSlot.Param.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	return &DataStoreRatingInitParamWithSlot{
		Slot:  types.NewPrimitiveS8(0),
		Param: NewDataStoreRatingInitParam(),
	}
}
