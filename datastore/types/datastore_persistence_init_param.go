// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePersistenceInitParam is sent in the PreparePostObject method
type DataStorePersistenceInitParam struct {
	types.Structure
	PersistenceSlotID *types.PrimitiveU16
	DeleteLastObject  *types.PrimitiveBool
}

// WriteTo writes the DataStorePersistenceInitParam to the given writable
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePersistenceInitParam.PersistenceSlotID.WriteTo(contentWritable)
	dataStorePersistenceInitParam.DeleteLastObject.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePersistenceInitParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePersistenceInitParam from the given readable
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePersistenceInitParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePersistenceInitParam header. %s", err.Error())
	}

	err = dataStorePersistenceInitParam.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.PersistenceSlotID. %s", err.Error())
	}

	err = dataStorePersistenceInitParam.DeleteLastObject.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.DeleteLastObject. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceInitParam
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Copy() types.RVType {
	copied := NewDataStorePersistenceInitParam()

	copied.StructureVersion = dataStorePersistenceInitParam.StructureVersion

	copied.PersistenceSlotID = dataStorePersistenceInitParam.PersistenceSlotID.Copy().(*types.PrimitiveU16)
	copied.DeleteLastObject = dataStorePersistenceInitParam.DeleteLastObject.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceInitParam); !ok {
		return false
	}

	other := o.(*DataStorePersistenceInitParam)

	if dataStorePersistenceInitParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePersistenceInitParam.PersistenceSlotID.Equals(other.PersistenceSlotID) {
		return false
	}

	if !dataStorePersistenceInitParam.DeleteLastObject.Equals(other.DeleteLastObject) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePersistenceInitParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s,\n", indentationValues, dataStorePersistenceInitParam.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDeleteLastObject: %s\n", indentationValues, dataStorePersistenceInitParam.DeleteLastObject))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInitParam returns a new DataStorePersistenceInitParam
func NewDataStorePersistenceInitParam() *DataStorePersistenceInitParam {
	return &DataStorePersistenceInitParam{
		PersistenceSlotID: types.NewPrimitiveU16(0),
		DeleteLastObject:  types.NewPrimitiveBool(false),
	}
}
