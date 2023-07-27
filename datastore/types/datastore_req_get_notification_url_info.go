// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqGetNotificationURLInfo is a data structure used by the DataStore protocol
type DataStoreReqGetNotificationURLInfo struct {
	nex.Structure
	URL        string
	Key        string
	Query      string
	RootCACert []byte
}

// ExtractFromStream extracts a DataStoreReqGetNotificationURLInfo structure from a stream
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqGetNotificationURLInfo.URL, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.URL. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.Key, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Key. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.Query, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Query. %s", err.Error())
	}

	dataStoreReqGetNotificationURLInfo.RootCACert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqGetNotificationURLInfo and returns a byte array
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetNotificationURLInfo.URL)
	stream.WriteString(dataStoreReqGetNotificationURLInfo.Key)
	stream.WriteString(dataStoreReqGetNotificationURLInfo.Query)
	stream.WriteBuffer(dataStoreReqGetNotificationURLInfo.RootCACert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetNotificationURLInfo
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetNotificationURLInfo()

	copied.URL = dataStoreReqGetNotificationURLInfo.URL
	copied.Key = dataStoreReqGetNotificationURLInfo.Key
	copied.Query = dataStoreReqGetNotificationURLInfo.Query
	copied.RootCACert = make([]byte, len(dataStoreReqGetNotificationURLInfo.RootCACert))

	copy(copied.RootCACert, dataStoreReqGetNotificationURLInfo.RootCACert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetNotificationURLInfo)

	if dataStoreReqGetNotificationURLInfo.URL != other.URL {
		return false
	}

	if dataStoreReqGetNotificationURLInfo.Key != other.Key {
		return false
	}

	if dataStoreReqGetNotificationURLInfo.Query != other.Query {
		return false
	}

	if !bytes.Equal(dataStoreReqGetNotificationURLInfo.RootCACert, other.RootCACert) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) String() string {
	return dataStoreReqGetNotificationURLInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetNotificationURLInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqGetNotificationURLInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, dataStoreReqGetNotificationURLInfo.URL))
	b.WriteString(fmt.Sprintf("%sKey: %q,\n", indentationValues, dataStoreReqGetNotificationURLInfo.Key))
	b.WriteString(fmt.Sprintf("%sQuery: %q,\n", indentationValues, dataStoreReqGetNotificationURLInfo.Query))
	b.WriteString(fmt.Sprintf("%sRootCaCert: %x\n", indentationValues, dataStoreReqGetNotificationURLInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetNotificationURLInfo returns a new DataStoreReqGetNotificationURLInfo
func NewDataStoreReqGetNotificationURLInfo() *DataStoreReqGetNotificationURLInfo {
	return &DataStoreReqGetNotificationURLInfo{}
}
