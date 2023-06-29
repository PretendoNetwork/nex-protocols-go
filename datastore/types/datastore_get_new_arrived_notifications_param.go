package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetNewArrivedNotificationsParam is a data structure used by the DataStore protocol
type DataStoreGetNewArrivedNotificationsParam struct {
	nex.Structure
	LastNotificationID uint64
	Limit              uint16
}

// ExtractFromStream extracts a DataStoreGetNewArrivedNotificationsParam structure from a stream
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetNewArrivedNotificationsParam.LastNotificationID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.LastNotificationID. %s", err.Error())
	}

	dataStoreGetNewArrivedNotificationsParam.Limit, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.Limit. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetNewArrivedNotificationsParam and returns a byte array
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetNewArrivedNotificationsParam.LastNotificationID)
	stream.WriteUInt16LE(dataStoreGetNewArrivedNotificationsParam.Limit)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetNewArrivedNotificationsParam
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetNewArrivedNotificationsParam()

	copied.LastNotificationID = dataStoreGetNewArrivedNotificationsParam.LastNotificationID
	copied.Limit = dataStoreGetNewArrivedNotificationsParam.Limit

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetNewArrivedNotificationsParam)

	if dataStoreGetNewArrivedNotificationsParam.LastNotificationID != other.LastNotificationID {
		return false
	}

	if dataStoreGetNewArrivedNotificationsParam.Limit != other.Limit {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) String() string {
	return dataStoreGetNewArrivedNotificationsParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetNewArrivedNotificationsParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLastNotificationID: %d,\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.LastNotificationID))
	b.WriteString(fmt.Sprintf("%sLimit: %d\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.Limit))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNewArrivedNotificationsParam returns a new DataStoreGetNewArrivedNotificationsParam
func NewDataStoreGetNewArrivedNotificationsParam() *DataStoreGetNewArrivedNotificationsParam {
	return &DataStoreGetNewArrivedNotificationsParam{}
}
