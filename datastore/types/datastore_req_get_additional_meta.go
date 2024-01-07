// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqGetAdditionalMeta is a data structure used by the DataStore protocol
type DataStoreReqGetAdditionalMeta struct {
	types.Structure
	OwnerID    *types.PID
	DataType   *types.PrimitiveU16
	Version    *types.PrimitiveU16
	MetaBinary *types.QBuffer
}

// ExtractFrom extracts the DataStoreReqGetAdditionalMeta from the given readable
func (d *DataStoreReqGetAdditionalMeta) ExtractFrom(readable types.Readable) error {
	var err error

	if err = d.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqGetAdditionalMeta header. %s", err.Error())
	}

	err = d.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.OwnerID. %s", err.Error())
	}

	err = d.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.DataType. %s", err.Error())
	}

	err = d.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.Version. %s", err.Error())
	}

	err = d.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.MetaBinary. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqGetAdditionalMeta to the given writable
func (d *DataStoreReqGetAdditionalMeta) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	d.OwnerID.WriteTo(contentWritable)
	d.DataType.WriteTo(contentWritable)
	d.Version.WriteTo(contentWritable)
	d.MetaBinary.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	d.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (d *DataStoreReqGetAdditionalMeta) Copy() types.RVType {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.StructureVersion = d.StructureVersion

	copied.OwnerID = d.OwnerID.Copy().(*types.PID)
	copied.DataType = d.DataType.Copy().(*types.PrimitiveU16)
	copied.Version = d.Version.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = d.MetaBinary.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (d *DataStoreReqGetAdditionalMeta) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetAdditionalMeta); !ok {
		return false
	}

	other := o.(*DataStoreReqGetAdditionalMeta)

	if d.StructureVersion != other.StructureVersion {
		return false
	}

	if !d.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !d.DataType.Equals(other.DataType) {
		return false
	}

	if !d.Version.Equals(other.Version) {
		return false
	}

	if !d.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (d *DataStoreReqGetAdditionalMeta) String() string {
	return d.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (d *DataStoreReqGetAdditionalMeta) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetAdditionalMeta{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, d.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, d.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, d.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, d.Version))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s\n", indentationValues, d.MetaBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	return &DataStoreReqGetAdditionalMeta{
		OwnerID:    types.NewPID(0),
		DataType:   types.NewPrimitiveU16(0),
		Version:    types.NewPrimitiveU16(0),
		MetaBinary: types.NewQBuffer(nil),
	}
}
