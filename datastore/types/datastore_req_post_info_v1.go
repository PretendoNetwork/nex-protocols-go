// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqPostInfoV1 is a data structure used by the DataStore protocol
type DataStoreReqPostInfoV1 struct {
	types.Structure
	DataID         *types.PrimitiveU32
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// ExtractFrom extracts the DataStoreReqPostInfoV1 from the given readable
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReqPostInfoV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqPostInfoV1 header. %s", err.Error())
	}

	err = dataStoreReqPostInfoV1.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.DataID. %s", err.Error())
	}

	err = dataStoreReqPostInfoV1.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.URL. %s", err.Error())
	}

	err = dataStoreReqPostInfoV1.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RequestHeaders. %s", err.Error())
	}

	err = dataStoreReqPostInfoV1.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.FormFields. %s", err.Error())
	}

	err = dataStoreReqPostInfoV1.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RootCACert. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqPostInfoV1 to the given writable
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReqPostInfoV1.DataID.WriteTo(contentWritable)
	dataStoreReqPostInfoV1.URL.WriteTo(contentWritable)
	dataStoreReqPostInfoV1.RequestHeaders.WriteTo(contentWritable)
	dataStoreReqPostInfoV1.FormFields.WriteTo(contentWritable)
	dataStoreReqPostInfoV1.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReqPostInfoV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqPostInfoV1
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Copy() types.RVType {
	copied := NewDataStoreReqPostInfoV1()

	copied.StructureVersion = dataStoreReqPostInfoV1.StructureVersion

	copied.DataID = dataStoreReqPostInfoV1.DataID.Copy().(*types.PrimitiveU32)
	copied.URL = dataStoreReqPostInfoV1.URL.Copy().(*types.String)
	copied.RequestHeaders = dataStoreReqPostInfoV1.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dataStoreReqPostInfoV1.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dataStoreReqPostInfoV1.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqPostInfoV1); !ok {
		return false
	}

	other := o.(*DataStoreReqPostInfoV1)

	if dataStoreReqPostInfoV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqPostInfoV1.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreReqPostInfoV1.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqPostInfoV1.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dataStoreReqPostInfoV1.FormFields.Equals(other.FormFields) {
		return false
	}

	return dataStoreReqPostInfoV1.RootCACert.Equals(other.RootCACert)
}

// String returns a string representation of the struct
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) String() string {
	return dataStoreReqPostInfoV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqPostInfoV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreReqPostInfoV1.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqPostInfoV1.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dataStoreReqPostInfoV1.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dataStoreReqPostInfoV1.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s\n", indentationValues, dataStoreReqPostInfoV1.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfoV1 returns a new DataStoreReqPostInfoV1
func NewDataStoreReqPostInfoV1() *DataStoreReqPostInfoV1 {
	dataStoreReqPostInfoV1 := &DataStoreReqPostInfoV1{
		DataID:         types.NewPrimitiveU32(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dataStoreReqPostInfoV1.RequestHeaders.Type = NewDataStoreKeyValue()
	dataStoreReqPostInfoV1.FormFields.Type = NewDataStoreKeyValue()

	return dataStoreReqPostInfoV1
}
