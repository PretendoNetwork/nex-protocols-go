// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreNotification is a data structure used by the DataStore protocol
type DataStoreNotification struct {
	types.Structure
	NotificationID *types.PrimitiveU64
	DataID         *types.PrimitiveU64
}

// ExtractFrom extracts the DataStoreNotification from the given readable
func (dataStoreNotification *DataStoreNotification) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreNotification.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreNotification header. %s", err.Error())
	}

	err = dataStoreNotification.NotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.NotificationID. %s", err.Error())
	}

	err = dataStoreNotification.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.DataID. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreNotification to the given writable
func (dataStoreNotification *DataStoreNotification) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreNotification.NotificationID.WriteTo(contentWritable)
	dataStoreNotification.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreNotification.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreNotification
func (dataStoreNotification *DataStoreNotification) Copy() types.RVType {
	copied := NewDataStoreNotification()

	copied.StructureVersion = dataStoreNotification.StructureVersion

	copied.NotificationID = dataStoreNotification.NotificationID.Copy().(*types.PrimitiveU64)
	copied.DataID = dataStoreNotification.DataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotification *DataStoreNotification) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreNotification); !ok {
		return false
	}

	other := o.(*DataStoreNotification)

	if dataStoreNotification.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreNotification.NotificationID.Equals(other.NotificationID) {
		return false
	}

	if !dataStoreNotification.DataID.Equals(other.DataID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreNotification.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNotificationID: %s,\n", indentationValues, dataStoreNotification.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %s\n", indentationValues, dataStoreNotification.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotification returns a new DataStoreNotification
func NewDataStoreNotification() *DataStoreNotification {
	return &DataStoreNotification{
		NotificationID: types.NewPrimitiveU64(0),
		DataID:         types.NewPrimitiveU64(0),
	}
}
