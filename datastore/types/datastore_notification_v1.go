// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreNotificationV1 is a data structure used by the DataStore protocol
type DataStoreNotificationV1 struct {
	types.Structure
	NotificationID *types.PrimitiveU64
	DataID         *types.PrimitiveU32
}

// ExtractFrom extracts the DataStoreNotificationV1 from the given readable
func (dataStoreNotificationV1 *DataStoreNotificationV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreNotificationV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreNotificationV1 header. %s", err.Error())
	}

	err = dataStoreNotificationV1.NotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.NotificationID. %s", err.Error())
	}

	err = dataStoreNotificationV1.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.DataID. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreNotificationV1 to the given writable
func (dataStoreNotificationV1 *DataStoreNotificationV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreNotificationV1.NotificationID.WriteTo(contentWritable)
	dataStoreNotificationV1.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreNotificationV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreNotificationV1
func (dataStoreNotificationV1 *DataStoreNotificationV1) Copy() types.RVType {
	copied := NewDataStoreNotificationV1()

	copied.StructureVersion = dataStoreNotificationV1.StructureVersion

	copied.NotificationID = dataStoreNotificationV1.NotificationID.Copy().(*types.PrimitiveU64)
	copied.DataID = dataStoreNotificationV1.DataID.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotificationV1 *DataStoreNotificationV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreNotificationV1); !ok {
		return false
	}

	other := o.(*DataStoreNotificationV1)

	if dataStoreNotificationV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreNotificationV1.NotificationID.Equals(other.NotificationID) {
		return false
	}

	if !dataStoreNotificationV1.DataID.Equals(other.DataID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreNotificationV1 *DataStoreNotificationV1) String() string {
	return dataStoreNotificationV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreNotificationV1 *DataStoreNotificationV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreNotificationV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreNotificationV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNotificationID: %s,\n", indentationValues, dataStoreNotificationV1.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %s\n", indentationValues, dataStoreNotificationV1.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotificationV1 returns a new DataStoreNotificationV1
func NewDataStoreNotificationV1() *DataStoreNotificationV1 {
	return &DataStoreNotificationV1{
		NotificationID: types.NewPrimitiveU64(0),
		DataID:         types.NewPrimitiveU32(0),
	}
}
