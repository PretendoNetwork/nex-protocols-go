// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetMetaParam is sent in the GetMeta method
type DataStoreGetMetaParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      *types.PrimitiveU8
	AccessPassword    *types.PrimitiveU64
}

// WriteTo writes the DataStoreGetMetaParam to the given writable
func (dataStoreGetMetaParam *DataStoreGetMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetMetaParam.DataID.WriteTo(contentWritable)
	dataStoreGetMetaParam.PersistenceTarget.WriteTo(contentWritable)
	dataStoreGetMetaParam.ResultOption.WriteTo(contentWritable)
	dataStoreGetMetaParam.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetMetaParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaParam from the given readable
func (dataStoreGetMetaParam *DataStoreGetMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetMetaParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetMetaParam header. %s", err.Error())
	}

	err = dataStoreGetMetaParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.DataID. %s", err.Error())
	}

	err = dataStoreGetMetaParam.PersistenceTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.PersistenceTarget. %s", err.Error())
	}

	err = dataStoreGetMetaParam.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.ResultOption. %s", err.Error())
	}

	err = dataStoreGetMetaParam.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetMetaParam
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaParam()

	copied.StructureVersion = dataStoreGetMetaParam.StructureVersion

	copied.DataID = dataStoreGetMetaParam.DataID.Copy().(*types.PrimitiveU64)

	copied.PersistenceTarget = dataStoreGetMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)

	copied.ResultOption = dataStoreGetMetaParam.ResultOption.Copy().(*types.PrimitiveU8)
	copied.AccessPassword = dataStoreGetMetaParam.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaParam)

	if dataStoreGetMetaParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetMetaParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreGetMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if !dataStoreGetMetaParam.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dataStoreGetMetaParam.AccessPassword.Equals(other.AccessPassword) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetMetaParam *DataStoreGetMetaParam) String() string {
	return dataStoreGetMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetMetaParam *DataStoreGetMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetMetaParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreGetMetaParam.DataID))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dataStoreGetMetaParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dataStoreGetMetaParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s\n", indentationValues, dataStoreGetMetaParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{
		DataID:            types.NewPrimitiveU64(0),
		PersistenceTarget: NewDataStorePersistenceTarget(),
		ResultOption:      types.NewPrimitiveU8(0),
		AccessPassword:    types.NewPrimitiveU64(0),
	}
}
