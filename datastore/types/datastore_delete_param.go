// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreDeleteParam is a data structure used by the DataStore protocol
type DataStoreDeleteParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	UpdatePassword *types.PrimitiveU64
}

// ExtractFrom extracts the DataStoreDeleteParam from the given readable
func (dataStoreDeleteParam *DataStoreDeleteParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreDeleteParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreDeleteParam header. %s", err.Error())
	}

	err = dataStoreDeleteParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.DataID. %s", err.Error())
	}

	err = dataStoreDeleteParam.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.UpdatePassword. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreDeleteParam to the given writable
func (dataStoreDeleteParam *DataStoreDeleteParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreDeleteParam.DataID.WriteTo(contentWritable)
	dataStoreDeleteParam.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreDeleteParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreDeleteParam *DataStoreDeleteParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaParamV1()

	copied.StructureVersion = dataStoreDeleteParam.StructureVersion

	copied.DataID = dataStoreDeleteParam.DataID.Copy().(*types.PrimitiveU64)
	copied.UpdatePassword = dataStoreDeleteParam.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreDeleteParam *DataStoreDeleteParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaParamV1); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaParamV1)

	if dataStoreDeleteParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreDeleteParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreDeleteParam.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreDeleteParam *DataStoreDeleteParam) String() string {
	return dataStoreDeleteParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreDeleteParam *DataStoreDeleteParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreDeleteParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreDeleteParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreDeleteParam.DataID))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s\n", indentationValues, dataStoreDeleteParam.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreDeleteParam returns a new DataStoreDeleteParam
func NewDataStoreDeleteParam() *DataStoreDeleteParam {
	return &DataStoreDeleteParam{
		DataID:         types.NewPrimitiveU64(0),
		UpdatePassword: types.NewPrimitiveU64(0),
	}
}
