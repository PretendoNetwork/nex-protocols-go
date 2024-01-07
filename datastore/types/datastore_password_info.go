// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePasswordInfo is a data structure used by the DataStore protocol
type DataStorePasswordInfo struct {
	types.Structure
	DataID         *types.PrimitiveU64
	AccessPassword *types.PrimitiveU64
	UpdatePassword *types.PrimitiveU64
}

// ExtractFrom extracts the DataStorePasswordInfo from the given readable
func (dataStorePasswordInfo *DataStorePasswordInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePasswordInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePasswordInfo header. %s", err.Error())
	}

	err = dataStorePasswordInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.DataID. %s", err.Error())
	}

	err = dataStorePasswordInfo.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.AccessPassword. %s", err.Error())
	}

	err = dataStorePasswordInfo.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.UpdatePassword. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePasswordInfo to the given writable
func (dataStorePasswordInfo *DataStorePasswordInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePasswordInfo.DataID.WriteTo(contentWritable)
	dataStorePasswordInfo.AccessPassword.WriteTo(contentWritable)
	dataStorePasswordInfo.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePasswordInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePasswordInfo
func (dataStorePasswordInfo *DataStorePasswordInfo) Copy() types.RVType {
	copied := NewDataStorePasswordInfo()

	copied.StructureVersion = dataStorePasswordInfo.StructureVersion

	copied.DataID = dataStorePasswordInfo.DataID.Copy().(*types.PrimitiveU64)
	copied.AccessPassword = dataStorePasswordInfo.AccessPassword.Copy().(*types.PrimitiveU64)
	copied.UpdatePassword = dataStorePasswordInfo.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePasswordInfo *DataStorePasswordInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePasswordInfo); !ok {
		return false
	}

	other := o.(*DataStorePasswordInfo)

	if dataStorePasswordInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePasswordInfo.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStorePasswordInfo.AccessPassword.Equals(other.AccessPassword) {
		return false
	}

	if !dataStorePasswordInfo.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePasswordInfo *DataStorePasswordInfo) String() string {
	return dataStorePasswordInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePasswordInfo *DataStorePasswordInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePasswordInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePasswordInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStorePasswordInfo.DataID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dataStorePasswordInfo.AccessPassword))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s\n", indentationValues, dataStorePasswordInfo.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePasswordInfo returns a new DataStorePasswordInfo
func NewDataStorePasswordInfo() *DataStorePasswordInfo {
	return &DataStorePasswordInfo{
		DataID:         types.NewPrimitiveU64(0),
		AccessPassword: types.NewPrimitiveU64(0),
		UpdatePassword: types.NewPrimitiveU64(0),
	}
}
