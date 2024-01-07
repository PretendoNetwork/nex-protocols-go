// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingTarget is sent in the RateObjects method
type DataStoreRatingTarget struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
}

// WriteTo writes the DataStoreRatingTarget to the given writable
func (dataStoreRatingTarget *DataStoreRatingTarget) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingTarget.DataID.WriteTo(contentWritable)
	dataStoreRatingTarget.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingTarget.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingTarget from the given readable
func (dataStoreRatingTarget *DataStoreRatingTarget) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingTarget.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingTarget header. %s", err.Error())
	}

	err = dataStoreRatingTarget.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.DataID. %s", err.Error())
	}

	err = dataStoreRatingTarget.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingTarget
func (dataStoreRatingTarget *DataStoreRatingTarget) Copy() types.RVType {
	copied := NewDataStoreRatingTarget()

	copied.StructureVersion = dataStoreRatingTarget.StructureVersion

	copied.DataID = dataStoreRatingTarget.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dataStoreRatingTarget.Slot.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingTarget *DataStoreRatingTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingTarget); !ok {
		return false
	}

	other := o.(*DataStoreRatingTarget)

	if dataStoreRatingTarget.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingTarget.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreRatingTarget.Slot.Equals(other.Slot) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRatingTarget *DataStoreRatingTarget) String() string {
	return dataStoreRatingTarget.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingTarget *DataStoreRatingTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingTarget{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingTarget.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreRatingTarget.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s\n", indentationValues, dataStoreRatingTarget.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingTarget returns a new DataStoreRatingTarget
func NewDataStoreRatingTarget() *DataStoreRatingTarget {
	return &DataStoreRatingTarget{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
	}
}
