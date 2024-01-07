// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreCompletePostParamV1 is a data structure used by the DataStore protocol
type DataStoreCompletePostParamV1 struct {
	types.Structure
	DataID    *types.PrimitiveU32
	IsSuccess *types.PrimitiveBool
}

// ExtractFrom extracts the DataStoreCompletePostParamV1 from the given readable
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCompletePostParamV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCompletePostParamV1 header. %s", err.Error())
	}

	err = dataStoreCompletePostParamV1.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.DataID. %s", err.Error())
	}

	err = dataStoreCompletePostParamV1.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.IsSuccess. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreCompletePostParamV1 to the given writable
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCompletePostParamV1.DataID.WriteTo(contentWritable)
	dataStoreCompletePostParamV1.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCompletePostParamV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCompletePostParamV1
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Copy() types.RVType {
	copied := NewDataStoreCompletePostParamV1()

	copied.StructureVersion = dataStoreCompletePostParamV1.StructureVersion

	copied.DataID = dataStoreCompletePostParamV1.DataID.Copy().(*types.PrimitiveU32)
	copied.IsSuccess = dataStoreCompletePostParamV1.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostParamV1); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostParamV1)

	if dataStoreCompletePostParamV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCompletePostParamV1.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreCompletePostParamV1.IsSuccess.Equals(other.IsSuccess) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) String() string {
	return dataStoreCompletePostParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCompletePostParamV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreCompletePostParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s\n", indentationValues, dataStoreCompletePostParamV1.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParamV1 returns a new DataStoreCompletePostParamV1
func NewDataStoreCompletePostParamV1() *DataStoreCompletePostParamV1 {
	return &DataStoreCompletePostParamV1{
		DataID:    types.NewPrimitiveU32(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}
}
