package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreKeyValue is sent in the PrepareGetObject method
type DataStoreKeyValue struct {
	nex.Structure
	Key   string
	Value string
}

// Bytes encodes the DataStoreKeyValue and returns a byte array
func (dataStoreKeyValue *DataStoreKeyValue) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreKeyValue.Key)
	stream.WriteString(dataStoreKeyValue.Value)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreKeyValue
func (dataStoreKeyValue *DataStoreKeyValue) Copy() nex.StructureInterface {
	copied := NewDataStoreKeyValue()

	copied.Key = dataStoreKeyValue.Key
	copied.Value = dataStoreKeyValue.Value

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreKeyValue *DataStoreKeyValue) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreKeyValue)

	if dataStoreKeyValue.Key != other.Key {
		return false
	}

	if dataStoreKeyValue.Value != other.Value {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreKeyValue.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sKey: %q,\n", indentationValues, dataStoreKeyValue.Key))
	b.WriteString(fmt.Sprintf("%sValue: %q\n", indentationValues, dataStoreKeyValue.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	return &DataStoreKeyValue{}
}
