// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetReplayMetaParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreGetReplayMetaParam struct {
	types.Structure
	ReplayID *types.PrimitiveU64
	MetaType *types.PrimitiveU8
}

// ExtractFrom extracts the DataStoreGetReplayMetaParam from the given readable
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetReplayMetaParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetReplayMetaParam header. %s", err.Error())
	}

	err = dataStoreGetReplayMetaParam.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.ReplayID. %s", err.Error())
	}

	err = dataStoreGetReplayMetaParam.MetaType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.MetaType. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetReplayMetaParam to the given writable
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetReplayMetaParam.ReplayID.WriteTo(contentWritable)
	dataStoreGetReplayMetaParam.MetaType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetReplayMetaParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetReplayMetaParam
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetReplayMetaParam()

	copied.StructureVersion = dataStoreGetReplayMetaParam.StructureVersion

	copied.ReplayID = dataStoreGetReplayMetaParam.ReplayID.Copy().(*types.PrimitiveU64)
	copied.MetaType = dataStoreGetReplayMetaParam.MetaType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetReplayMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetReplayMetaParam)

	if dataStoreGetReplayMetaParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetReplayMetaParam.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dataStoreGetReplayMetaParam.MetaType.Equals(other.MetaType) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) String() string {
	return dataStoreGetReplayMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetReplayMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetReplayMetaParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dataStoreGetReplayMetaParam.ReplayID))
	b.WriteString(fmt.Sprintf("%sMetaType: %s\n", indentationValues, dataStoreGetReplayMetaParam.MetaType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetReplayMetaParam returns a new DataStoreGetReplayMetaParam
func NewDataStoreGetReplayMetaParam() *DataStoreGetReplayMetaParam {
	return &DataStoreGetReplayMetaParam{
		ReplayID: types.NewPrimitiveU64(0),
		MetaType: types.NewPrimitiveU8(0),
	}
}
