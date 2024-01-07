// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetNotificationURLParam is a data structure used by the DataStore protocol
type DataStoreGetNotificationURLParam struct {
	types.Structure
	PreviousURL *types.String
}

// ExtractFrom extracts the DataStoreGetNotificationURLParam from the given readable
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetNotificationURLParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetNotificationURLParam header. %s", err.Error())
	}

	err = dataStoreGetNotificationURLParam.PreviousURL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNotificationURLParam.PreviousURL. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetNotificationURLParam to the given writable
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetNotificationURLParam.PreviousURL.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetNotificationURLParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetNotificationURLParam
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Copy() types.RVType {
	copied := NewDataStoreGetNotificationURLParam()

	copied.StructureVersion = dataStoreGetNotificationURLParam.StructureVersion

	copied.PreviousURL = dataStoreGetNotificationURLParam.PreviousURL

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNotificationURLParam *DataStoreGetNotificationURLParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetNotificationURLParam); !ok {
		return false
	}

	other := o.(*DataStoreGetNotificationURLParam)

	if dataStoreGetNotificationURLParam.StructureVersion != other.StructureVersion {
		return false
	}

	return dataStoreGetNotificationURLParam.PreviousURL.Equals(other.PreviousURL)
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetNotificationURLParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPreviousURL: %s\n", indentationValues, dataStoreGetNotificationURLParam.PreviousURL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNotificationURLParam returns a new DataStoreGetNotificationURLParam
func NewDataStoreGetNotificationURLParam() *DataStoreGetNotificationURLParam {
	return &DataStoreGetNotificationURLParam{
		PreviousURL: types.NewString(""),
	}
}
