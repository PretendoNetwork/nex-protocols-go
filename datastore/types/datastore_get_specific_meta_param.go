// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetSpecificMetaParam is a data structure used by the DataStore protocol
type DataStoreGetSpecificMetaParam struct {
	types.Structure
	DataIDs *types.List[*types.PrimitiveU64]
}

// ExtractFrom extracts the DataStoreGetSpecificMetaParam from the given readable
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetSpecificMetaParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetSpecificMetaParam header. %s", err.Error())
	}

	err = dataStoreGetSpecificMetaParam.DataIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParam.DataIDs. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetSpecificMetaParam to the given writable
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetSpecificMetaParam.DataIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetSpecificMetaParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParam
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetSpecificMetaParam()

	copied.StructureVersion = dataStoreGetSpecificMetaParam.StructureVersion

	copied.DataIDs = dataStoreGetSpecificMetaParam.DataIDs.Copy().(*types.List[*types.PrimitiveU64])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetSpecificMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetSpecificMetaParam)

	if dataStoreGetSpecificMetaParam.StructureVersion != other.StructureVersion {
		return false
	}

	if dataStoreGetSpecificMetaParam.DataIDs.Equals(other.DataIDs) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) String() string {
	return dataStoreGetSpecificMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetSpecificMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetSpecificMetaParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataIDs: %s\n", indentationValues, dataStoreGetSpecificMetaParam.DataIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetSpecificMetaParam returns a new DataStoreGetSpecificMetaParam
func NewDataStoreGetSpecificMetaParam() *DataStoreGetSpecificMetaParam {
	dataStoreGetSpecificMetaParam := &DataStoreGetSpecificMetaParam{
		DataIDs: types.NewList[*types.PrimitiveU64](),
	}

	dataStoreGetSpecificMetaParam.DataIDs.Type = types.NewPrimitiveU64(0)

	return dataStoreGetSpecificMetaParam
}
