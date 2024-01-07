// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	types.Structure
	DataID  *types.PrimitiveU64
	GetInfo *datastore_types.DataStoreReqGetInfo
}

// ExtractFrom extracts the DataStoreFileServerObjectInfo from the given readable
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreFileServerObjectInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreFileServerObjectInfo header. %s", err.Error())
	}

	err = dataStoreFileServerObjectInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.DataID. %s", err.Error())
	}

	err = dataStoreFileServerObjectInfo.GetInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.GetInfo. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreFileServerObjectInfo to the given writable
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreFileServerObjectInfo.DataID.WriteTo(contentWritable)
	dataStoreFileServerObjectInfo.GetInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreFileServerObjectInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreFileServerObjectInfo
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Copy() types.RVType {
	copied := NewDataStoreFileServerObjectInfo()

	copied.StructureVersion = dataStoreFileServerObjectInfo.StructureVersion

	copied.DataID = dataStoreFileServerObjectInfo.DataID.Copy().(*types.PrimitiveU64)
	copied.GetInfo = dataStoreFileServerObjectInfo.GetInfo.Copy().(*datastore_types.DataStoreReqGetInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFileServerObjectInfo); !ok {
		return false
	}

	other := o.(*DataStoreFileServerObjectInfo)

	if dataStoreFileServerObjectInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreFileServerObjectInfo.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreFileServerObjectInfo.GetInfo.Equals(other.GetInfo) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) String() string {
	return dataStoreFileServerObjectInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFileServerObjectInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreFileServerObjectInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreFileServerObjectInfo.DataID))
	b.WriteString(fmt.Sprintf("%sGetInfo: %s\n", indentationValues, dataStoreFileServerObjectInfo.GetInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{
		DataID: types.NewPrimitiveU64(0),
		GetInfo: datastore_types.NewDataStoreReqGetInfo(),
	}
}
