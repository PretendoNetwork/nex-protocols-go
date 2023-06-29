package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreNotification is a data structure used by the DataStore protocol
type DataStoreNotification struct {
	nex.Structure
	NotificationID uint64
	DataID         uint64
}

// ExtractFromStream extracts a DataStoreNotification structure from a stream
func (dataStoreNotification *DataStoreNotification) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreNotification.NotificationID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.NotificationID. %s", err.Error())
	}

	dataStoreNotification.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.DataID. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreNotification and returns a byte array
func (dataStoreNotification *DataStoreNotification) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreNotification.NotificationID)
	stream.WriteUInt64LE(dataStoreNotification.DataID)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreNotification
func (dataStoreNotification *DataStoreNotification) Copy() nex.StructureInterface {
	copied := NewDataStoreNotification()

	copied.NotificationID = dataStoreNotification.NotificationID
	copied.DataID = dataStoreNotification.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotification *DataStoreNotification) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreNotification)

	if dataStoreNotification.NotificationID != other.NotificationID {
		return false
	}

	if dataStoreNotification.DataID != other.DataID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreNotification *DataStoreNotification) String() string {
	return dataStoreNotification.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreNotification *DataStoreNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreNotification{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreNotification.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sNotificationID: %d,\n", indentationValues, dataStoreNotification.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %d\n", indentationValues, dataStoreNotification.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotification returns a new DataStoreNotification
func NewDataStoreNotification() *DataStoreNotification {
	return &DataStoreNotification{}
}
