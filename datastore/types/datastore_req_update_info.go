// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqUpdateInfo is a data structure used by the DataStore protocol
type DataStoreReqUpdateInfo struct {
	types.Structure
	Version        *types.PrimitiveU32
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// ExtractFrom extracts the DataStoreReqUpdateInfo from the given readable
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStoreReqUpdateInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqUpdateInfo header. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		version, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreReqUpdateInfo.Version.Value = version
	} else {
		version, err := readable.ReadPrimitiveUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreReqUpdateInfo.Version.Value = *types.PrimitiveU32(version)
	}

	err = dataStoreReqUpdateInfo.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.URL. %s", err.Error())
	}

	err = dataStoreReqUpdateInfo.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RequestHeaders. %s", err.Error())
	}

	err = dataStoreReqUpdateInfo.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.FormFields. %s", err.Error())
	}

	err = dataStoreReqUpdateInfo.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqUpdateInfo to the given writable
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	contentWritable := writable.CopyNew()

	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		contentWritable.WritePrimitiveUInt32LE(dataStoreReqUpdateInfo.Version.Value)
	} else {
		contentWritable.WritePrimitiveUInt16LE(*types.PrimitiveU16(dataStoreReqUpdateInfo.Version.Value))
	}

	dataStoreReqUpdateInfo.URL.WriteTo(contentWritable)
	dataStoreReqUpdateInfo.RequestHeaders.WriteTo(contentWritable)
	dataStoreReqUpdateInfo.FormFields.WriteTo(contentWritable)
	dataStoreReqUpdateInfo.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReqUpdateInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqUpdateInfo
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Copy() types.RVType {
	copied := NewDataStoreReqUpdateInfo()

	copied.StructureVersion = dataStoreReqUpdateInfo.StructureVersion

	copied.Version = dataStoreReqUpdateInfo.Version.Copy().(*types.PrimitiveU32)
	copied.URL = dataStoreReqUpdateInfo.URL.Copy().(*types.String)
	copied.RequestHeaders = dataStoreReqUpdateInfo.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dataStoreReqUpdateInfo.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dataStoreReqUpdateInfo.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqUpdateInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqUpdateInfo)

	if dataStoreReqUpdateInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqUpdateInfo.Version.Equals(other.Version) {
		return false
	}

	if !dataStoreReqUpdateInfo.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqUpdateInfo.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dataStoreReqUpdateInfo.FormFields.Equals(other.FormFields) {
		return false
	}

	return dataStoreReqUpdateInfo.RootCACert.Equals(other.RootCACert)
}

// String returns a string representation of the struct
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) String() string {
	return dataStoreReqUpdateInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqUpdateInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqUpdateInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dataStoreReqUpdateInfo.Version))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqUpdateInfo.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dataStoreReqUpdateInfo.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dataStoreReqUpdateInfo.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s\n", indentationValues, dataStoreReqUpdateInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqUpdateInfo returns a new DataStoreReqUpdateInfo
func NewDataStoreReqUpdateInfo() *DataStoreReqUpdateInfo {
	dataStoreReqUpdateInfo := &DataStoreReqUpdateInfo{
		Version:        types.NewPrimitiveU32(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dataStoreReqUpdateInfo.RequestHeaders.Type = NewDataStoreKeyValue()
	dataStoreReqUpdateInfo.FormFields.Type = NewDataStoreKeyValue()

	return dataStoreReqUpdateInfo
}
