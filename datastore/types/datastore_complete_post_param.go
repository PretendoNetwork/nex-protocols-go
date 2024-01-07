// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreCompletePostParam is sent in the CompletePostObject method
type DataStoreCompletePostParam struct {
	types.Structure
	DataID    *types.PrimitiveU64
	IsSuccess *types.PrimitiveBool
}

// WriteTo writes the DataStoreCompletePostParam to the given writable
func (dataStoreCompletePostParam *DataStoreCompletePostParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCompletePostParam.DataID.WriteTo(contentWritable)
	dataStoreCompletePostParam.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCompletePostParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompletePostParam from the given readable
func (dataStoreCompletePostParam *DataStoreCompletePostParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCompletePostParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCompletePostParam header. %s", err.Error())
	}

	err = dataStoreCompletePostParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.DataID. %s", err.Error())
	}

	err = dataStoreCompletePostParam.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostParam
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostParam()

	copied.StructureVersion = dataStoreCompletePostParam.StructureVersion

	copied.DataID = dataStoreCompletePostParam.DataID.Copy().(*types.PrimitiveU64)
	copied.IsSuccess = dataStoreCompletePostParam.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostParam); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostParam)

	if dataStoreCompletePostParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCompletePostParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreCompletePostParam.IsSuccess.Equals(other.IsSuccess) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostParam *DataStoreCompletePostParam) String() string {
	return dataStoreCompletePostParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostParam *DataStoreCompletePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCompletePostParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreCompletePostParam.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s\n", indentationValues, dataStoreCompletePostParam.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParam returns a new DataStoreCompletePostParam
func NewDataStoreCompletePostParam() *DataStoreCompletePostParam {
	return &DataStoreCompletePostParam{
		DataID:    types.NewPrimitiveU64(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}
}
