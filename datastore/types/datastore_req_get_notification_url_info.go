package datastore_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreReqGetNotificationURLInfo struct {
	nex.Structure
	Url        string
	Key        string
	Query      string
	RootCaCert []byte
}

// ExtractFromStream extracts a DataStoreReqGetNotificationURLInfo structure from a stream
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqGetNotificationURLInfo.Url, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Url. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.Key, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Key. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.Query, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Query. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.RootCaCert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.RootCaCert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqGetNotificationURLInfo and returns a byte array
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetNotificationURLInfo.Url)
	stream.WriteString(dataStoreReqGetNotificationURLInfo.Key)
	stream.WriteString(dataStoreReqGetNotificationURLInfo.Query)
	stream.WriteBuffer(dataStoreReqGetNotificationURLInfo.RootCaCert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetNotificationURLInfo
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetNotificationURLInfo()

	copied.Url = dataStoreReqGetNotificationURLInfo.Url
	copied.Key = dataStoreReqGetNotificationURLInfo.Key
	copied.Query = dataStoreReqGetNotificationURLInfo.Query
	copied.RootCaCert = make([]byte, len(dataStoreReqGetNotificationURLInfo.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqGetNotificationURLInfo.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetNotificationURLInfo)

	if dataStoreReqGetNotificationURLInfo.Url != other.Url {
		return false
	}

	if dataStoreReqGetNotificationURLInfo.Key != other.Key {
		return false
	}

	if dataStoreReqGetNotificationURLInfo.Query != other.Query {
		return false
	}

	if !bytes.Equal(dataStoreReqGetNotificationURLInfo.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
}

// NewDataStoreReqGetNotificationURLInfo returns a new DataStoreReqGetNotificationURLInfo
func NewDataStoreReqGetNotificationURLInfo() *DataStoreReqGetNotificationURLInfo {
	return &DataStoreReqGetNotificationURLInfo{}
}
