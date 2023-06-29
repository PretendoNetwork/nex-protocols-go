package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePermission contains information about a permission for a DataStore object
type DataStorePermission struct {
	nex.Structure
	Permission   uint8
	RecipientIds []uint32
}

// ExtractFromStream extracts a DataStorePermission structure from a stream
func (dataStorePermission *DataStorePermission) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePermission.Permission, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.Permission. %s", err.Error())
	}

	dataStorePermission.RecipientIds, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.RecipientIds. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePermission and returns a byte array
func (dataStorePermission *DataStorePermission) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePermission.Permission)
	stream.WriteListUInt32LE(dataStorePermission.RecipientIds)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePermission
func (dataStorePermission *DataStorePermission) Copy() nex.StructureInterface {
	copied := NewDataStorePermission()

	copied.Permission = dataStorePermission.Permission
	copied.RecipientIds = make([]uint32, len(dataStorePermission.RecipientIds))

	copy(copied.RecipientIds, dataStorePermission.RecipientIds)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePermission *DataStorePermission) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePermission)

	if dataStorePermission.Permission != other.Permission {
		return false
	}

	if len(dataStorePermission.RecipientIds) != len(other.RecipientIds) {
		return false
	}

	for i := 0; i < len(dataStorePermission.RecipientIds); i++ {
		if dataStorePermission.RecipientIds[i] != other.RecipientIds[i] {
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
	b.WriteString(fmt.Sprintf("%sRecipientIds: %v\n", indentationValues, dataStorePermission.RecipientIds))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	return &DataStorePermission{}
}
