// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreCustomRankingRatingCondition holds data for the DataStore (Super Mario Maker) protocol
type DataStoreCustomRankingRatingCondition struct {
	nex.Structure
	Slot     int8
	MinValue int32
	MaxValue int32
}

// ExtractFromStream extracts a DataStoreCustomRankingRatingCondition structure from a stream
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCustomRankingRatingCondition.Slot, err = stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.Slot from stream. %s", err.Error())
	}

	dataStoreCustomRankingRatingCondition.MinValue, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MinValue from stream. %s", err.Error())
	}

	dataStoreCustomRankingRatingCondition.MaxValue, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxValue from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCustomRankingRatingCondition and returns a byte array
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteInt8(dataStoreCustomRankingRatingCondition.Slot)
	stream.WriteInt32LE(dataStoreCustomRankingRatingCondition.MinValue)
	stream.WriteInt32LE(dataStoreCustomRankingRatingCondition.MaxValue)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCustomRankingRatingCondition
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) Copy() nex.StructureInterface {
	copied := NewDataStoreCustomRankingRatingCondition()

	copied.Slot = dataStoreCustomRankingRatingCondition.Slot
	copied.MinValue = dataStoreCustomRankingRatingCondition.MinValue
	copied.MaxValue = dataStoreCustomRankingRatingCondition.MaxValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingRatingCondition *DataStoreCustomRankingRatingCondition) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCustomRankingRatingCondition)

	if dataStoreCustomRankingRatingCondition.Slot != other.Slot {
		return false
	}

	if dataStoreCustomRankingRatingCondition.MinValue != other.MinValue {
		return false
	}

	if dataStoreCustomRankingRatingCondition.MaxValue != other.MaxValue {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCustomRankingRatingCondition.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, dataStoreCustomRankingRatingCondition.Slot))
	b.WriteString(fmt.Sprintf("%sMinValue: %d,\n", indentationValues, dataStoreCustomRankingRatingCondition.MinValue))
	b.WriteString(fmt.Sprintf("%sMaxValue: %d,\n", indentationValues, dataStoreCustomRankingRatingCondition.MaxValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingRatingCondition returns a new DataStoreCustomRankingRatingCondition
func NewDataStoreCustomRankingRatingCondition() *DataStoreCustomRankingRatingCondition {
	return &DataStoreCustomRankingRatingCondition{}
}
