// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreNotificationV1 is a data structure used by the DataStore protocol
type DataStoreNotificationV1 struct {
	nex.Structure
	NotificationID uint64
	DataID         uint32
}

// ExtractFromStream extracts a DataStoreNotificationV1 structure from a stream
func (dataStoreNotificationV1 *DataStoreNotificationV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreNotificationV1.NotificationID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.NotificationID. %s", err.Error())
	}

	dataStoreNotificationV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.DataID. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreNotificationV1 and returns a byte array
func (dataStoreNotificationV1 *DataStoreNotificationV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreNotificationV1.NotificationID)
	stream.WriteUInt32LE(dataStoreNotificationV1.DataID)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreNotificationV1
func (dataStoreNotificationV1 *DataStoreNotificationV1) Copy() nex.StructureInterface {
	copied := NewDataStoreNotificationV1()

	copied.NotificationID = dataStoreNotificationV1.NotificationID
	copied.DataID = dataStoreNotificationV1.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotificationV1 *DataStoreNotificationV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreNotificationV1)

	if dataStoreNotificationV1.NotificationID != other.NotificationID {
		return false
	}

	if dataStoreNotificationV1.DataID != other.DataID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreNotificationV1 *DataStoreNotificationV1) String() string {
	return dataStoreNotificationV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreNotificationV1 *DataStoreNotificationV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreNotificationV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreNotificationV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sNotificationID: %d,\n", indentationValues, dataStoreNotificationV1.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %d\n", indentationValues, dataStoreNotificationV1.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotificationV1 returns a new DataStoreNotificationV1
func NewDataStoreNotificationV1() *DataStoreNotificationV1 {
	return &DataStoreNotificationV1{}
}
