// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	nex.Structure
	OwnerID           uint32
	PersistenceSlotID uint16
}

// ExtractFromStream extracts a DataStorePersistenceTarget structure from a stream
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceTarget.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.OwnerID. %s", err.Error())
	}

	dataStorePersistenceTarget.PersistenceSlotID, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.PersistenceSlotID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of dataStorePersistenceTarget
func (dataStorePersistenceTarget *DataStorePersistenceTarget) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceTarget()

	copied.SetStructureVersion(dataStorePersistenceTarget.StructureVersion())

	copied.OwnerID = dataStorePersistenceTarget.OwnerID
	copied.PersistenceSlotID = dataStorePersistenceTarget.PersistenceSlotID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceTarget *DataStorePersistenceTarget) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceTarget)

	if dataStorePersistenceTarget.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStorePersistenceTarget.OwnerID != other.OwnerID {
		return false
	}

	if dataStorePersistenceTarget.PersistenceSlotID != other.PersistenceSlotID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePersistenceTarget *DataStorePersistenceTarget) String() string {
	return dataStorePersistenceTarget.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePersistenceTarget *DataStorePersistenceTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceTarget{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePersistenceTarget.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStorePersistenceTarget.OwnerID))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %d\n", indentationValues, dataStorePersistenceTarget.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{
		OwnerID:           0,
		PersistenceSlotID: 0,
	}
}
