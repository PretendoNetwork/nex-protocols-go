// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqGetNotificationURLInfo is a data structure used by the DataStore protocol
type DataStoreReqGetNotificationURLInfo struct {
	types.Structure
	URL        *types.String
	Key        *types.String
	Query      *types.String
	RootCACert *types.Buffer
}

// ExtractFrom extracts the DataStoreReqGetNotificationURLInfo from the given readable
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReqGetNotificationURLInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqGetNotificationURLInfo header. %s", err.Error())
	}

	err = dataStoreReqGetNotificationURLInfo.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.URL. %s", err.Error())
	}

	err = dataStoreReqGetNotificationURLInfo.Key.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Key. %s", err.Error())
	}

	err = dataStoreReqGetNotificationURLInfo.Query.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Query. %s", err.Error())
	}

	err = dataStoreReqGetNotificationURLInfo.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqGetNotificationURLInfo to the given writable
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReqGetNotificationURLInfo.URL.WriteTo(contentWritable)
	dataStoreReqGetNotificationURLInfo.Key.WriteTo(contentWritable)
	dataStoreReqGetNotificationURLInfo.Query.WriteTo(contentWritable)
	dataStoreReqGetNotificationURLInfo.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReqGetNotificationURLInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqGetNotificationURLInfo
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Copy() types.RVType {
	copied := NewDataStoreReqGetNotificationURLInfo()

	copied.StructureVersion = dataStoreReqGetNotificationURLInfo.StructureVersion

	copied.URL = dataStoreReqGetNotificationURLInfo.URL.Copy().(*types.String)
	copied.Key = dataStoreReqGetNotificationURLInfo.Key.Copy().(*types.String)
	copied.Query = dataStoreReqGetNotificationURLInfo.Query.Copy().(*types.String)
	copied.RootCACert = dataStoreReqGetNotificationURLInfo.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetNotificationURLInfo *DataStoreReqGetNotificationURLInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetNotificationURLInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqGetNotificationURLInfo)

	if dataStoreReqGetNotificationURLInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqGetNotificationURLInfo.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqGetNotificationURLInfo.Key.Equals(other.Key) {
		return false
	}

	if !dataStoreReqGetNotificationURLInfo.Query.Equals(other.Query) {
		return false
	}

	if !dataStoreReqGetNotificationURLInfo.RootCACert.Equals(other.RootCACert) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqGetNotificationURLInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqGetNotificationURLInfo.URL))
	b.WriteString(fmt.Sprintf("%sKey: %s,\n", indentationValues, dataStoreReqGetNotificationURLInfo.Key))
	b.WriteString(fmt.Sprintf("%sQuery: %s,\n", indentationValues, dataStoreReqGetNotificationURLInfo.Query))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s\n", indentationValues, dataStoreReqGetNotificationURLInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetNotificationURLInfo returns a new DataStoreReqGetNotificationURLInfo
func NewDataStoreReqGetNotificationURLInfo() *DataStoreReqGetNotificationURLInfo {
	return &DataStoreReqGetNotificationURLInfo{
		URL:        types.NewString(""),
		Key:        types.NewString(""),
		Query:      types.NewString(""),
		RootCACert: types.NewBuffer(nil),
	}
}
