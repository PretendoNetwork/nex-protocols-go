// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetSpecificMetaParamV1 is a data structure used by the DataStore protocol
type DataStoreGetSpecificMetaParamV1 struct {
	nex.Structure
	DataIDs []uint32
}

// ExtractFromStream extracts a DataStoreGetSpecificMetaParamV1 structure from a stream
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetSpecificMetaParamV1.DataIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParamV1.DataIDs. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreGetSpecificMetaParamV1 and returns a byte array
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetSpecificMetaParamV1.DataIDs)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParamV1
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreGetSpecificMetaParamV1()

	copied.DataIDs = make([]uint32, len(dataStoreGetSpecificMetaParamV1.DataIDs))

	copy(copied.DataIDs, dataStoreGetSpecificMetaParamV1.DataIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetSpecificMetaParamV1)

	if len(dataStoreGetSpecificMetaParamV1.DataIDs) != len(other.DataIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetSpecificMetaParamV1.DataIDs); i++ {
		if dataStoreGetSpecificMetaParamV1.DataIDs[i] != other.DataIDs[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) String() string {
	return dataStoreGetSpecificMetaParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetSpecificMetaParamV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetSpecificMetaParamV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataIDs: %v\n", indentationValues, dataStoreGetSpecificMetaParamV1.DataIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetSpecificMetaParamV1 returns a new DataStoreGetSpecificMetaParamV1
func NewDataStoreGetSpecificMetaParamV1() *DataStoreGetSpecificMetaParamV1 {
	return &DataStoreGetSpecificMetaParamV1{}
}
