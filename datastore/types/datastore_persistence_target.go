// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	types.Structure
	OwnerID           *types.PID
	PersistenceSlotID *types.PrimitiveU16
}

// WriteTo writes the DataStorePersistenceTarget to the given writable
func (dataStorePersistenceTarget *DataStorePersistenceTarget) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePersistenceTarget.OwnerID.WriteTo(contentWritable)
	dataStorePersistenceTarget.PersistenceSlotID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePersistenceTarget.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePersistenceTarget from the given readable
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePersistenceTarget.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePersistenceTarget header. %s", err.Error())
	}

	err = dataStorePersistenceTarget.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.OwnerID. %s", err.Error())
	}

	err = dataStorePersistenceTarget.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.PersistenceSlotID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of dataStorePersistenceTarget
func (dataStorePersistenceTarget *DataStorePersistenceTarget) Copy() types.RVType {
	copied := NewDataStorePersistenceTarget()

	copied.StructureVersion = dataStorePersistenceTarget.StructureVersion

	copied.OwnerID = dataStorePersistenceTarget.OwnerID.Copy().(*types.PID)
	copied.PersistenceSlotID = dataStorePersistenceTarget.PersistenceSlotID.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceTarget *DataStorePersistenceTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceTarget); !ok {
		return false
	}

	other := o.(*DataStorePersistenceTarget)

	if dataStorePersistenceTarget.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePersistenceTarget.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dataStorePersistenceTarget.PersistenceSlotID.Equals(other.PersistenceSlotID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePersistenceTarget.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dataStorePersistenceTarget.OwnerID))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s\n", indentationValues, dataStorePersistenceTarget.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{
		OwnerID:           types.NewPID(0),
		PersistenceSlotID: types.NewPrimitiveU16(0),
	}
}
