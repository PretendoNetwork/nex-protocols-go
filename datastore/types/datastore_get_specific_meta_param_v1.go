// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetSpecificMetaParamV1 is a data structure used by the DataStore protocol
type DataStoreGetSpecificMetaParamV1 struct {
	types.Structure
	DataIDs *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the DataStoreGetSpecificMetaParamV1 from the given readable
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetSpecificMetaParamV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetSpecificMetaParamV1 header. %s", err.Error())
	}

	err = dataStoreGetSpecificMetaParamV1.DataIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParamV1.DataIDs. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetSpecificMetaParamV1 to the given writable
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetSpecificMetaParamV1.DataIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetSpecificMetaParamV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParamV1
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Copy() types.RVType {
	copied := NewDataStoreGetSpecificMetaParamV1()

	copied.StructureVersion = dataStoreGetSpecificMetaParamV1.StructureVersion

	copied.DataIDs = dataStoreGetSpecificMetaParamV1.DataIDs.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetSpecificMetaParamV1); !ok {
		return false
	}

	other := o.(*DataStoreGetSpecificMetaParamV1)

	if dataStoreGetSpecificMetaParamV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetSpecificMetaParamV1.DataIDs.Equals(other.DataIDs) {
		return false
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetSpecificMetaParamV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataIDs: %s\n", indentationValues, dataStoreGetSpecificMetaParamV1.DataIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetSpecificMetaParamV1 returns a new DataStoreGetSpecificMetaParamV1
func NewDataStoreGetSpecificMetaParamV1() *DataStoreGetSpecificMetaParamV1 {
	dataStoreGetSpecificMetaParamV1 := &DataStoreGetSpecificMetaParamV1{
		DataIDs: types.NewList[*types.PrimitiveU32](),
	}

	dataStoreGetSpecificMetaParamV1.DataIDs.Type = types.NewPrimitiveU32(0)

	return dataStoreGetSpecificMetaParamV1
}
