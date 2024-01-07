// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePermission contains information about a permission for a DataStore object
type DataStorePermission struct {
	types.Structure
	Permission   *types.PrimitiveU8
	RecipientIDs *types.List[*types.PID]
}

// ExtractFrom extracts the DataStorePermission from the given readable
func (dataStorePermission *DataStorePermission) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePermission.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePermission header. %s", err.Error())
	}

	err = dataStorePermission.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.Permission. %s", err.Error())
	}

	err = dataStorePermission.RecipientIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.RecipientIDs. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePermission to the given writable
func (dataStorePermission *DataStorePermission) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePermission.Permission.WriteTo(contentWritable)
	dataStorePermission.RecipientIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePermission.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePermission
func (dataStorePermission *DataStorePermission) Copy() types.RVType {
	copied := NewDataStorePermission()

	copied.StructureVersion = dataStorePermission.StructureVersion

	copied.Permission = dataStorePermission.Permission.Copy().(*types.PrimitiveU8)
	copied.RecipientIDs = dataStorePermission.RecipientIDs.Copy().(*types.List[*types.PID])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePermission *DataStorePermission) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePermission); !ok {
		return false
	}

	other := o.(*DataStorePermission)

	if dataStorePermission.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePermission.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePermission.RecipientIDs.Equals(other.RecipientIDs) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePermission *DataStorePermission) String() string {
	return dataStorePermission.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePermission *DataStorePermission) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePermission{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePermission.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStorePermission.Permission))
	b.WriteString(fmt.Sprintf("%sRecipientIDs: %s\n", indentationValues, dataStorePermission.RecipientIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	dataStorePermission := &DataStorePermission{
		Permission:   types.NewPrimitiveU8(0),
		RecipientIDs: types.NewList[*types.PID](),
	}

	dataStorePermission.RecipientIDs.Type = types.NewPID(0)

	return dataStorePermission
}
