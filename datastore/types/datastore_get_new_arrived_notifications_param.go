package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

// NewDataStoreGetNewArrivedNotificationsParam returns a new DataStoreGetNewArrivedNotificationsParam
func NewDataStoreGetNewArrivedNotificationsParam() *DataStoreGetNewArrivedNotificationsParam {
	return &DataStoreGetNewArrivedNotificationsParam{}
}
