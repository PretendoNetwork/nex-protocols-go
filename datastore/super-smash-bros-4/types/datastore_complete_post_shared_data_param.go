// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCompletePostSharedDataParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreCompletePostSharedDataParam struct {
	types.Structure
	DataID        *types.PrimitiveU64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// ExtractFrom extracts the DataStoreCompletePostSharedDataParam from the given readable
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCompletePostSharedDataParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCompletePostSharedDataParam header. %s", err.Error())
	}

	err = dataStoreCompletePostSharedDataParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.DataID. %s", err.Error())
	}

	err = dataStoreCompletePostSharedDataParam.CompleteParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.CompleteParam. %s", err.Error())
	}

	err = dataStoreCompletePostSharedDataParam.PrepareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.PrepareParam. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreCompletePostSharedDataParam to the given writable
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCompletePostSharedDataParam.DataID.WriteTo(contentWritable)
	dataStoreCompletePostSharedDataParam.CompleteParam.WriteTo(contentWritable)
	dataStoreCompletePostSharedDataParam.PrepareParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCompletePostSharedDataParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCompletePostSharedDataParam
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostSharedDataParam()

	copied.StructureVersion = dataStoreCompletePostSharedDataParam.StructureVersion

	copied.DataID = dataStoreCompletePostSharedDataParam.DataID.Copy().(*types.PrimitiveU64)
	copied.CompleteParam = dataStoreCompletePostSharedDataParam.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostSharedDataParam.PrepareParam.Copy().(*DataStorePreparePostSharedDataParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostSharedDataParam)

	if dataStoreCompletePostSharedDataParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) String() string {
	return dataStoreCompletePostSharedDataParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCompletePostSharedDataParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreCompletePostSharedDataParam.DataID))
	b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dataStoreCompletePostSharedDataParam.CompleteParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareParam: %s\n", indentationValues, dataStoreCompletePostSharedDataParam.PrepareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostSharedDataParam returns a new DataStoreCompletePostSharedDataParam
func NewDataStoreCompletePostSharedDataParam() *DataStoreCompletePostSharedDataParam {
	return &DataStoreCompletePostSharedDataParam{
		DataID: types.NewPrimitiveU64(0),
		CompleteParam: datastore_types.NewDataStoreCompletePostParam(),
		PrepareParam: NewDataStorePreparePostSharedDataParam(),
	}
}
