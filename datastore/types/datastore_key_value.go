package datastore_types

import "github.com/PretendoNetwork/nex-go"

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

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	return &DataStoreKeyValue{}
}
