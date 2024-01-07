// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetNewArrivedNotificationsParam is a data structure used by the DataStore protocol
type DataStoreGetNewArrivedNotificationsParam struct {
	types.Structure
	LastNotificationID *types.PrimitiveU64
	Limit              *types.PrimitiveU16
}

// ExtractFrom extracts the DataStoreGetNewArrivedNotificationsParam from the given readable
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetNewArrivedNotificationsParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetNewArrivedNotificationsParam header. %s", err.Error())
	}

	err = dataStoreGetNewArrivedNotificationsParam.LastNotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.LastNotificationID. %s", err.Error())
	}

	err = dataStoreGetNewArrivedNotificationsParam.Limit.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.Limit. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetNewArrivedNotificationsParam to the given writable
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetNewArrivedNotificationsParam.LastNotificationID.WriteTo(contentWritable)
	dataStoreGetNewArrivedNotificationsParam.Limit.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetNewArrivedNotificationsParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetNewArrivedNotificationsParam
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Copy() types.RVType {
	copied := NewDataStoreGetNewArrivedNotificationsParam()

	copied.StructureVersion = dataStoreGetNewArrivedNotificationsParam.StructureVersion

	copied.LastNotificationID = dataStoreGetNewArrivedNotificationsParam.LastNotificationID
	copied.Limit = dataStoreGetNewArrivedNotificationsParam.Limit

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetNewArrivedNotificationsParam); !ok {
		return false
	}

	other := o.(*DataStoreGetNewArrivedNotificationsParam)

	if dataStoreGetNewArrivedNotificationsParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetNewArrivedNotificationsParam.LastNotificationID.Equals(other.LastNotificationID) {
		return false
	}

	if !dataStoreGetNewArrivedNotificationsParam.Limit.Equals(other.Limit) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLastNotificationID: %s,\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.LastNotificationID))
	b.WriteString(fmt.Sprintf("%sLimit: %s\n", indentationValues, dataStoreGetNewArrivedNotificationsParam.Limit))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNewArrivedNotificationsParam returns a new DataStoreGetNewArrivedNotificationsParam
func NewDataStoreGetNewArrivedNotificationsParam() *DataStoreGetNewArrivedNotificationsParam {
	return &DataStoreGetNewArrivedNotificationsParam{
		LastNotificationID: types.NewPrimitiveU64(0),
		Limit:              types.NewPrimitiveU16(0),
	}
}
