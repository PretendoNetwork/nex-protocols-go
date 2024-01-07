// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqGetAdditionalMeta is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreReqGetAdditionalMeta struct {
	types.Structure
	OwnerID    *types.PrimitiveU32
	DataType   *types.PrimitiveU16
	Version    *types.PrimitiveU16
	MetaBinary []byte
}

// ExtractFrom extracts the DataStoreReqGetAdditionalMeta from the given readable
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReqGetAdditionalMeta.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqGetAdditionalMeta header. %s", err.Error())
	}

	err = dataStoreReqGetAdditionalMeta.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.OwnerID. %s", err.Error())
	}

	err = dataStoreReqGetAdditionalMeta.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.DataType. %s", err.Error())
	}

	err = dataStoreReqGetAdditionalMeta.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.Version. %s", err.Error())
	}

	dataStoreReqGetAdditionalMeta.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.MetaBinary. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqGetAdditionalMeta to the given writable
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReqGetAdditionalMeta.OwnerID.WriteTo(contentWritable)
	dataStoreReqGetAdditionalMeta.DataType.WriteTo(contentWritable)
	dataStoreReqGetAdditionalMeta.Version.WriteTo(contentWritable)
	stream.WriteQBuffer(dataStoreReqGetAdditionalMeta.MetaBinary)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Copy() types.RVType {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.StructureVersion = dataStoreReqGetAdditionalMeta.StructureVersion

	copied.OwnerID = dataStoreReqGetAdditionalMeta.OwnerID
	copied.DataType = dataStoreReqGetAdditionalMeta.DataType
	copied.Version = dataStoreReqGetAdditionalMeta.Version
	copied.MetaBinary = make([]byte, len(dataStoreReqGetAdditionalMeta.MetaBinary))

	copy(copied.MetaBinary, dataStoreReqGetAdditionalMeta.MetaBinary)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetAdditionalMeta); !ok {
		return false
	}

	other := o.(*DataStoreReqGetAdditionalMeta)

	if dataStoreReqGetAdditionalMeta.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqGetAdditionalMeta.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dataStoreReqGetAdditionalMeta.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreReqGetAdditionalMeta.Version.Equals(other.Version) {
		return false
	}

	if !dataStoreReqGetAdditionalMeta.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) String() string {
	return dataStoreReqGetAdditionalMeta.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetAdditionalMeta{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, dataStoreReqGetAdditionalMeta.Version))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x\n", indentationValues, dataStoreReqGetAdditionalMeta.MetaBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	return &DataStoreReqGetAdditionalMeta{}
}
