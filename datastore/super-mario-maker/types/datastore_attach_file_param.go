// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreAttachFileParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreAttachFileParam struct {
	types.Structure
	PostParam   *datastore_types.DataStorePreparePostParam
	ReferDataID *types.PrimitiveU64
	ContentType *types.String
}

// ExtractFrom extracts the DataStoreAttachFileParam from the given readable
func (dataStoreAttachFileParam *DataStoreAttachFileParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreAttachFileParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreAttachFileParam header. %s", err.Error())
	}

	err = dataStoreAttachFileParam.PostParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.PostParam from stream. %s", err.Error())
	}

	err = dataStoreAttachFileParam.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ReferDataID from stream. %s", err.Error())
	}

	err = dataStoreAttachFileParam.ContentType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ContentType from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreAttachFileParam to the given writable
func (dataStoreAttachFileParam *DataStoreAttachFileParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreAttachFileParam.PostParam.WriteTo(contentWritable)
	dataStoreAttachFileParam.ReferDataID.WriteTo(contentWritable)
	dataStoreAttachFileParam.ContentType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreAttachFileParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreAttachFileParam
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Copy() types.RVType {
	copied := NewDataStoreAttachFileParam()

	copied.StructureVersion = dataStoreAttachFileParam.StructureVersion

	copied.PostParam = dataStoreAttachFileParam.PostParam.Copy().(*datastore_types.DataStorePreparePostParam)
	copied.ReferDataID = dataStoreAttachFileParam.ReferDataID.Copy().(*types.PrimitiveU64)
	copied.ContentType = dataStoreAttachFileParam.ContentType.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreAttachFileParam); !ok {
		return false
	}

	other := o.(*DataStoreAttachFileParam)

	if dataStoreAttachFileParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreAttachFileParam.PostParam.Equals(other.PostParam) {
		return false
	}

	if !dataStoreAttachFileParam.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dataStoreAttachFileParam.ContentType.Equals(other.ContentType) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreAttachFileParam *DataStoreAttachFileParam) String() string {
	return dataStoreAttachFileParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreAttachFileParam *DataStoreAttachFileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreAttachFileParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreAttachFileParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPostParam: %s\n", indentationValues, dataStoreAttachFileParam.PostParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dataStoreAttachFileParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sContentType: %s,\n", indentationValues, dataStoreAttachFileParam.ContentType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreAttachFileParam returns a new DataStoreAttachFileParam
func NewDataStoreAttachFileParam() *DataStoreAttachFileParam {
	return &DataStoreAttachFileParam{
		PostParam:   datastore_types.NewDataStorePreparePostParam(),
		ReferDataID: types.NewPrimitiveU64(0),
		ContentType: types.NewString(""),
	}
}
