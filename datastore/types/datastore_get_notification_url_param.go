package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreGetNotificationURLParam struct {
	nex.Structure
	PreviousUrl string
}

// ExtractFromStream extracts a DataStoreGetNotificationURLParam structure from a stream
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetNotificationURLParam.PreviousUrl, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNotificationURLParam.PreviousUrl. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetNotificationURLParam and returns a byte array
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreGetNotificationURLParam.PreviousUrl)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetNotificationURLParam
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetNotificationURLParam()

	copied.PreviousUrl = dataStoreGetNotificationURLParam.PreviousUrl

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetNotificationURLParam)

	return dataStoreGetNotificationURLParam.PreviousUrl != other.PreviousUrl
}

// NewDataStoreGetNotificationURLParam returns a new DataStoreGetNotificationURLParam
func NewDataStoreGetNotificationURLParam() *DataStoreGetNotificationURLParam {
	return &DataStoreGetNotificationURLParam{}
}
