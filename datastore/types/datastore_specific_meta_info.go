// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSpecificMetaInfo is a type within the DataStore protocol
type DataStoreSpecificMetaInfo struct {
	types.Structure
	DataID   types.UInt64
	OwnerID  types.PID
	Size     types.UInt32
	DataType types.UInt16
	Version  types.UInt32
}

// WriteTo writes the DataStoreSpecificMetaInfo to the given writable
func (dssmi DataStoreSpecificMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssmi.DataID.WriteTo(contentWritable)
	dssmi.OwnerID.WriteTo(contentWritable)
	dssmi.Size.WriteTo(contentWritable)
	dssmi.DataType.WriteTo(contentWritable)
	dssmi.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssmi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSpecificMetaInfo from the given readable
func (dssmi *DataStoreSpecificMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssmi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo header. %s", err.Error())
	}

	err = dssmi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataID. %s", err.Error())
	}

	err = dssmi.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.OwnerID. %s", err.Error())
	}

	err = dssmi.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Size. %s", err.Error())
	}

	err = dssmi.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataType. %s", err.Error())
	}

	err = dssmi.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Version. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfo
func (dssmi DataStoreSpecificMetaInfo) Copy() types.RVType {
	copied := NewDataStoreSpecificMetaInfo()

	copied.StructureVersion = dssmi.StructureVersion
	copied.DataID = dssmi.DataID.Copy().(types.UInt64)
	copied.OwnerID = dssmi.OwnerID.Copy().(types.PID)
	copied.Size = dssmi.Size.Copy().(types.UInt32)
	copied.DataType = dssmi.DataType.Copy().(types.UInt16)
	copied.Version = dssmi.Version.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given DataStoreSpecificMetaInfo contains the same data as the current DataStoreSpecificMetaInfo
func (dssmi DataStoreSpecificMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreSpecificMetaInfo); !ok {
		return false
	}

	other := o.(DataStoreSpecificMetaInfo)

	if dssmi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssmi.DataID.Equals(other.DataID) {
		return false
	}

	if !dssmi.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dssmi.Size.Equals(other.Size) {
		return false
	}

	if !dssmi.DataType.Equals(other.DataType) {
		return false
	}

	return dssmi.Version.Equals(other.Version)
}

// CopyRef copies the current value of the DataStoreSpecificMetaInfo
// and returns a pointer to the new copy
func (dssmi DataStoreSpecificMetaInfo) CopyRef() types.RVTypePtr {
	copied := dssmi.Copy().(DataStoreSpecificMetaInfo)
	return &copied
}

// Deref takes a pointer to the DataStoreSpecificMetaInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dssmi *DataStoreSpecificMetaInfo) Deref() types.RVType {
	return *dssmi
}

// String returns the string representation of the DataStoreSpecificMetaInfo
func (dssmi DataStoreSpecificMetaInfo) String() string {
	return dssmi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSpecificMetaInfo using the provided indentation level
func (dssmi DataStoreSpecificMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dssmi.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dssmi.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dssmi.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dssmi.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dssmi.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfo returns a new DataStoreSpecificMetaInfo
func NewDataStoreSpecificMetaInfo() DataStoreSpecificMetaInfo {
	return DataStoreSpecificMetaInfo{
		DataID:   types.NewUInt64(0),
		OwnerID:  types.NewPID(0),
		Size:     types.NewUInt32(0),
		DataType: types.NewUInt16(0),
		Version:  types.NewUInt32(0),
	}

}
