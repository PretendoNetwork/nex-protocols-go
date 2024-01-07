// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqPostInfo is sent in the PreparePostObject method
type DataStoreReqPostInfo struct {
	types.Structure
	DataID         *types.PrimitiveU64
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// WriteTo writes the DataStoreReqPostInfo to the given writable
func (dataStoreReqPostInfo *DataStoreReqPostInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReqPostInfo.DataID.WriteTo(contentWritable)
	dataStoreReqPostInfo.URL.WriteTo(contentWritable)
	dataStoreReqPostInfo.RequestHeaders.WriteTo(contentWritable)
	dataStoreReqPostInfo.FormFields.WriteTo(contentWritable)
	dataStoreReqPostInfo.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReqPostInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqPostInfo from the given readable
func (dataStoreReqPostInfo *DataStoreReqPostInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReqPostInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqPostInfo header. %s", err.Error())
	}

	err = dataStoreReqPostInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.DataID. %s", err.Error())
	}

	err = dataStoreReqPostInfo.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.URL. %s", err.Error())
	}

	err = dataStoreReqPostInfo.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.RequestHeaders. %s", err.Error())
	}

	err = dataStoreReqPostInfo.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.FormFields. %s", err.Error())
	}

	err = dataStoreReqPostInfo.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqPostInfo
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Copy() types.RVType {
	copied := NewDataStoreReqPostInfo()

	copied.StructureVersion = dataStoreReqPostInfo.StructureVersion

	copied.DataID = dataStoreReqPostInfo.DataID.Copy().(*types.PrimitiveU64)
	copied.URL = dataStoreReqPostInfo.URL.Copy().(*types.String)
	copied.RequestHeaders = dataStoreReqPostInfo.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dataStoreReqPostInfo.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dataStoreReqPostInfo.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqPostInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqPostInfo)

	if dataStoreReqPostInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqPostInfo.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreReqPostInfo.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqPostInfo.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dataStoreReqPostInfo.FormFields.Equals(other.FormFields) {
		return false
	}

	return dataStoreReqPostInfo.RootCACert.Equals(other.RootCACert)
}

// String returns a string representation of the struct
func (dataStoreReqPostInfo *DataStoreReqPostInfo) String() string {
	return dataStoreReqPostInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqPostInfo *DataStoreReqPostInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqPostInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreReqPostInfo.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqPostInfo.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dataStoreReqPostInfo.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dataStoreReqPostInfo.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s\n", indentationValues, dataStoreReqPostInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfo returns a new DataStoreReqPostInfo
func NewDataStoreReqPostInfo() *DataStoreReqPostInfo {
	dataStoreReqPostInfo := &DataStoreReqPostInfo{
		DataID:         types.NewPrimitiveU64(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dataStoreReqPostInfo.RequestHeaders.Type = NewDataStoreKeyValue()
	dataStoreReqPostInfo.FormFields.Type = NewDataStoreKeyValue()

	return dataStoreReqPostInfo
}
