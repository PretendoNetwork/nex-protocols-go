// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSpecificMetaInfo is a data structure used by the DataStore protocol
type DataStoreSpecificMetaInfo struct {
	types.Structure
	DataID   *types.PrimitiveU64
	OwnerID  *types.PID
	Size     *types.PrimitiveU32
	DataType *types.PrimitiveU16
	Version  *types.PrimitiveU32
}

// ExtractFrom extracts the DataStoreSpecificMetaInfo from the given readable
func (d *DataStoreSpecificMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = d.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSpecificMetaInfo header. %s", err.Error())
	}

	err = d.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataID. %s", err.Error())
	}

	err = d.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.OwnerID. %s", err.Error())
	}

	err = d.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Size. %s", err.Error())
	}

	err = d.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataType. %s", err.Error())
	}

	err = d.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Version. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreSpecificMetaInfo to the given writable
func (d *DataStoreSpecificMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	d.DataID.WriteTo(contentWritable)
	d.OwnerID.WriteTo(contentWritable)
	d.Size.WriteTo(contentWritable)
	d.DataType.WriteTo(contentWritable)
	d.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	d.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfo
func (d *DataStoreSpecificMetaInfo) Copy() types.RVType {
	copied := NewDataStoreSpecificMetaInfo()

	copied.StructureVersion = d.StructureVersion

	copied.DataID = d.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = d.OwnerID.Copy().(*types.PID)
	copied.Size = d.Size.Copy().(*types.PrimitiveU32)
	copied.DataType = d.DataType.Copy().(*types.PrimitiveU16)
	copied.Version = d.Version.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (d *DataStoreSpecificMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSpecificMetaInfo); !ok {
		return false
	}

	other := o.(*DataStoreSpecificMetaInfo)

	if d.StructureVersion != other.StructureVersion {
		return false
	}

	if !d.DataID.Equals(other.DataID) {
		return false
	}

	if !d.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !d.Size.Equals(other.Size) {
		return false
	}

	if !d.DataType.Equals(other.DataType) {
		return false
	}

	if !d.Version.Equals(other.Version) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (d *DataStoreSpecificMetaInfo) String() string {
	return d.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (d *DataStoreSpecificMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, d.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, d.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, d.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, d.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, d.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %s\n", indentationValues, d.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfo returns a new DataStoreSpecificMetaInfo
func NewDataStoreSpecificMetaInfo() *DataStoreSpecificMetaInfo {
	return &DataStoreSpecificMetaInfo{
		DataID:   types.NewPrimitiveU64(0),
		OwnerID:  types.NewPID(0),
		Size:     types.NewPrimitiveU32(0),
		DataType: types.NewPrimitiveU16(0),
		Version:  types.NewPrimitiveU32(0),
	}
}
