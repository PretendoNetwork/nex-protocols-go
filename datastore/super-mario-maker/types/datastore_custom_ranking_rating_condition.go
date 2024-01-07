// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreCustomRankingRatingCondition holds data for the DataStore (Super Mario Maker) protocol
type DataStoreCustomRankingRatingCondition struct {
	types.Structure
	Slot     *types.PrimitiveS8
	MinValue *types.PrimitiveS32
	MaxValue *types.PrimitiveS32
	MinCount *types.PrimitiveS32 // * Revision 1
	MaxCount *types.PrimitiveS32 // * Revision 1
}

// ExtractFrom extracts the DataStoreCustomRankingRatingCondition from the given readable
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCustomRankingRatingCondition.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCustomRankingRatingCondition header. %s", err.Error())
	}

	err = dataStoreCustomRankingRatingCondition.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.Slot from stream. %s", err.Error())
	}

	err = dataStoreCustomRankingRatingCondition.MinValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MinValue from stream. %s", err.Error())
	}

	err = dataStoreCustomRankingRatingCondition.MaxValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxValue from stream. %s", err.Error())
	}

	if dataStoreCustomRankingRatingCondition.StructureVersion >= 1 {
		err = dataStoreCustomRankingRatingCondition.MaxCount.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxCount from stream. %s", err.Error())
		}

		err = dataStoreCustomRankingRatingCondition.MaxCount.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxCount from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the DataStoreCustomRankingRatingCondition to the given writable
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCustomRankingRatingCondition.Slot.WriteTo(contentWritable)
	dataStoreCustomRankingRatingCondition.MinValue.WriteTo(contentWritable)
	dataStoreCustomRankingRatingCondition.MaxValue.WriteTo(contentWritable)

	if dataStoreCustomRankingRatingCondition.StructureVersion >= 1 {
		dataStoreCustomRankingRatingCondition.MinCount.WriteTo(contentWritable)
		dataStoreCustomRankingRatingCondition.MaxCount.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStoreCustomRankingRatingCondition.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCustomRankingRatingCondition
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) Copy() types.RVType {
	copied := NewDataStoreCustomRankingRatingCondition()

	copied.StructureVersion = dataStoreCustomRankingRatingCondition.StructureVersion

	copied.Slot = dataStoreCustomRankingRatingCondition.Slot.Copy().(*types.PrimitiveS8)
	copied.MinValue = dataStoreCustomRankingRatingCondition.MinValue.Copy().(*types.PrimitiveS32)
	copied.MaxValue = dataStoreCustomRankingRatingCondition.MaxValue.Copy().(*types.PrimitiveS32)
	copied.MinCount = dataStoreCustomRankingRatingCondition.MinCount.Copy().(*types.PrimitiveS32)
	copied.MaxCount = dataStoreCustomRankingRatingCondition.MaxCount.Copy().(*types.PrimitiveS32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCustomRankingRatingCondition); !ok {
		return false
	}

	other := o.(*DataStoreCustomRankingRatingCondition)

	if dataStoreCustomRankingRatingCondition.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCustomRankingRatingCondition.Slot.Equals(other.Slot) {
		return false
	}

	if !dataStoreCustomRankingRatingCondition.MinValue.Equals(other.MinValue) {
		return false
	}

	if !dataStoreCustomRankingRatingCondition.MaxValue.Equals(other.MaxValue) {
		return false
	}

	if !dataStoreCustomRankingRatingCondition.MinCount.Equals(other.MinCount) {
		return false
	}

	if !dataStoreCustomRankingRatingCondition.MaxCount.Equals(other.MaxCount) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) String() string {
	return dataStoreCustomRankingRatingCondition.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCustomRankingRatingCondition{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCustomRankingRatingCondition.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dataStoreCustomRankingRatingCondition.Slot))
	b.WriteString(fmt.Sprintf("%sMinValue: %s,\n", indentationValues, dataStoreCustomRankingRatingCondition.MinValue))
	b.WriteString(fmt.Sprintf("%sMaxValue: %s,\n", indentationValues, dataStoreCustomRankingRatingCondition.MaxValue))

	if dataStoreCustomRankingRatingCondition.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sMinCount: %s,\n", indentationValues, dataStoreCustomRankingRatingCondition.MinCount))
		b.WriteString(fmt.Sprintf("%sMaxCount: %s\n", indentationValues, dataStoreCustomRankingRatingCondition.MaxCount))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingRatingCondition returns a new DataStoreCustomRankingRatingCondition
func NewDataStoreCustomRankingRatingCondition() *DataStoreCustomRankingRatingCondition {
	return &DataStoreCustomRankingRatingCondition{
		Slot:     types.NewPrimitiveS8(0),
		MinValue: types.NewPrimitiveS32(0),
		MaxValue: types.NewPrimitiveS32(0),
		MinCount: types.NewPrimitiveS32(0),
		MaxCount: types.NewPrimitiveS32(0),
	}
}
