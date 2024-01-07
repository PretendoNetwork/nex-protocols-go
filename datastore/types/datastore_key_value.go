// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreKeyValue is sent in the PrepareGetObject method
type DataStoreKeyValue struct {
	types.Structure
	Key   *types.String
	Value *types.String
}

// WriteTo writes the DataStoreKeyValue to the given writable
func (dataStoreKeyValue *DataStoreKeyValue) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreKeyValue.Key.WriteTo(contentWritable)
	dataStoreKeyValue.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreKeyValue.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreKeyValue from the given readable
func (dataStoreKeyValue *DataStoreKeyValue) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreKeyValue.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreKeyValue header. %s", err.Error())
	}

	err = dataStoreKeyValue.Key.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreKeyValue.Key. %s", err.Error())
	}

	err = dataStoreKeyValue.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreKeyValue.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreKeyValue
func (dataStoreKeyValue *DataStoreKeyValue) Copy() types.RVType {
	copied := NewDataStoreKeyValue()

	copied.StructureVersion = dataStoreKeyValue.StructureVersion

	copied.Key = dataStoreKeyValue.Key.Copy().(*types.String)
	copied.Value = dataStoreKeyValue.Value.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreKeyValue *DataStoreKeyValue) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreKeyValue); !ok {
		return false
	}

	other := o.(*DataStoreKeyValue)

	if dataStoreKeyValue.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreKeyValue.Key.Equals(other.Key) {
		return false
	}

	if !dataStoreKeyValue.Value.Equals(other.Value) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreKeyValue *DataStoreKeyValue) String() string {
	return dataStoreKeyValue.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreKeyValue *DataStoreKeyValue) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreKeyValue{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreKeyValue.StructureVersion))
	b.WriteString(fmt.Sprintf("%sKey: %s,\n", indentationValues, dataStoreKeyValue.Key))
	b.WriteString(fmt.Sprintf("%sValue: %s\n", indentationValues, dataStoreKeyValue.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	return &DataStoreKeyValue{
		Key:   types.NewString(""),
		Value: types.NewString(""),
	}
}
