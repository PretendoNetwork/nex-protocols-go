// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetNotificationURLParam is a data structure used by the DataStore protocol
type DataStoreGetNotificationURLParam struct {
	nex.Structure
	PreviousURL string
}

// ExtractFromStream extracts a DataStoreGetNotificationURLParam structure from a stream
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetNotificationURLParam.PreviousURL, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNotificationURLParam.PreviousURL. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetNotificationURLParam and returns a byte array
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreGetNotificationURLParam.PreviousURL)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetNotificationURLParam
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetNotificationURLParam()

	copied.SetStructureVersion(dataStoreGetNotificationURLParam.StructureVersion())

	copied.PreviousURL = dataStoreGetNotificationURLParam.PreviousURL

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetNotificationURLParam)

	if dataStoreGetNotificationURLParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	return dataStoreGetNotificationURLParam.PreviousURL == other.PreviousURL
}

// String returns a string representation of the struct
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) String() string {
	return dataStoreGetNotificationURLParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetNotificationURLParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetNotificationURLParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPreviousURL: %q\n", indentationValues, dataStoreGetNotificationURLParam.PreviousURL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNotificationURLParam returns a new DataStoreGetNotificationURLParam
func NewDataStoreGetNotificationURLParam() *DataStoreGetNotificationURLParam {
	return &DataStoreGetNotificationURLParam{}
}
