// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePermission contains information about a permission for a DataStore object
type DataStorePermission struct {
	nex.Structure
	Permission   uint8
	RecipientIDs []uint32
}

// ExtractFromStream extracts a DataStorePermission structure from a stream
func (dataStorePermission *DataStorePermission) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePermission.Permission, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.Permission. %s", err.Error())
	}

	dataStorePermission.RecipientIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.RecipientIDs. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePermission and returns a byte array
func (dataStorePermission *DataStorePermission) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePermission.Permission)
	stream.WriteListUInt32LE(dataStorePermission.RecipientIDs)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePermission
func (dataStorePermission *DataStorePermission) Copy() nex.StructureInterface {
	copied := NewDataStorePermission()

	copied.SetStructureVersion(dataStorePermission.StructureVersion())

	copied.Permission = dataStorePermission.Permission
	copied.RecipientIDs = make([]uint32, len(dataStorePermission.RecipientIDs))

	copy(copied.RecipientIDs, dataStorePermission.RecipientIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePermission *DataStorePermission) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePermission)

	if dataStorePermission.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStorePermission.Permission != other.Permission {
		return false
	}

	if len(dataStorePermission.RecipientIDs) != len(other.RecipientIDs) {
		return false
	}

	for i := 0; i < len(dataStorePermission.RecipientIDs); i++ {
		if dataStorePermission.RecipientIDs[i] != other.RecipientIDs[i] {
			return false
		}
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePermission.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPermission: %d,\n", indentationValues, dataStorePermission.Permission))
	b.WriteString(fmt.Sprintf("%sRecipientIDs: %v\n", indentationValues, dataStorePermission.RecipientIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	return &DataStorePermission{
		Permission:   0,
		RecipientIDs: make([]uint32, 0),
	}
}
