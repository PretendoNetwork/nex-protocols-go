// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePersistenceInitParam is sent in the PreparePostObject method
type DataStorePersistenceInitParam struct {
	nex.Structure
	PersistenceSlotID uint16
	DeleteLastObject  bool
}

// ExtractFromStream extracts a DataStorePersistenceInitParam structure from a stream
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceInitParam.PersistenceSlotID, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.PersistenceSlotID. %s", err.Error())
	}

	dataStorePersistenceInitParam.DeleteLastObject, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.DeleteLastObject. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceInitParam
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInitParam()

	copied.PersistenceSlotID = dataStorePersistenceInitParam.PersistenceSlotID
	copied.DeleteLastObject = dataStorePersistenceInitParam.DeleteLastObject

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInitParam)

	if dataStorePersistenceInitParam.PersistenceSlotID != other.PersistenceSlotID {
		return false
	}

	if dataStorePersistenceInitParam.DeleteLastObject != other.DeleteLastObject {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) String() string {
	return dataStorePersistenceInitParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceInitParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePersistenceInitParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %d,\n", indentationValues, dataStorePersistenceInitParam.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDeleteLastObject: %t\n", indentationValues, dataStorePersistenceInitParam.DeleteLastObject))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInitParam returns a new DataStorePersistenceInitParam
func NewDataStorePersistenceInitParam() *DataStorePersistenceInitParam {
	return &DataStorePersistenceInitParam{}
}
