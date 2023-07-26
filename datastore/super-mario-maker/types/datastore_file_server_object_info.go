// Package types implements all the types used by the DataStore Super Mario Maker protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	nex.Structure
	DataID  uint64
	GetInfo *datastore_types.DataStoreReqGetInfo
}

// Bytes encodes the DataStoreFileServerObjectInfo and returns a byte array
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreFileServerObjectInfo.DataID)
	stream.WriteStructure(dataStoreFileServerObjectInfo.GetInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFileServerObjectInfo
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreFileServerObjectInfo()

	copied.DataID = dataStoreFileServerObjectInfo.DataID
	copied.GetInfo = dataStoreFileServerObjectInfo.GetInfo.Copy().(*datastore_types.DataStoreReqGetInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFileServerObjectInfo)

	if dataStoreFileServerObjectInfo.DataID != other.DataID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreFileServerObjectInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreFileServerObjectInfo.DataID))

	if dataStoreFileServerObjectInfo.GetInfo != nil {
		b.WriteString(fmt.Sprintf("%sGetInfo: %s\n", indentationValues, dataStoreFileServerObjectInfo.GetInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGetInfo: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{}
}
