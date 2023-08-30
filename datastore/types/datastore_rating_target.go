// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

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

	copied.SetStructureVersion(dataStoreRatingTarget.StructureVersion())

	copied.DataID = dataStoreRatingTarget.DataID
	copied.Slot = dataStoreRatingTarget.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingTarget *DataStoreRatingTarget) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingTarget)

	if dataStoreRatingTarget.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreRatingTarget.DataID != other.DataID {
		return false
	}

	if dataStoreRatingTarget.Slot != other.Slot {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRatingTarget.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreRatingTarget.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %d\n", indentationValues, dataStoreRatingTarget.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingTarget returns a new DataStoreRatingTarget
func NewDataStoreRatingTarget() *DataStoreRatingTarget {
	return &DataStoreRatingTarget{}
}
