// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqGetInfo is sent in the PrepareGetObject method
type DataStoreReqGetInfo struct {
	types.Structure
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	Size           *types.PrimitiveU32
	RootCACert     *types.Buffer
	DataID         *types.PrimitiveU64             // NEX 3.5.0+
}

// WriteTo writes the DataStoreReqGetInfo to the given writable
func (dataStoreReqGetInfo *DataStoreReqGetInfo) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	contentWritable := writable.CopyNew()


	dataStoreReqGetInfo.URL.WriteTo(contentWritable)
	dataStoreReqGetInfo.RequestHeaders.WriteTo(contentWritable)
	dataStoreReqGetInfo.Size.WriteTo(contentWritable)
	dataStoreReqGetInfo.RootCACert.WriteTo(contentWritable)

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStoreReqGetInfo.DataID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStoreReqGetInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqGetInfo from the given readable
func (dataStoreReqGetInfo *DataStoreReqGetInfo) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStoreReqGetInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePreparePostParam header. %s", err.Error())
	}

	err = dataStoreReqGetInfo.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.URL. %s", err.Error())
	}

	err = dataStoreReqGetInfo.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.RequestHeaders. %s", err.Error())
	}

	err = dataStoreReqGetInfo.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Size. %s", err.Error())
	}

	err = dataStoreReqGetInfo.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.RootCACert. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
	err = 	dataStoreReqGetInfo.DataID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePreparePostParam.DataID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqGetInfo
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Copy() types.RVType {
	copied := NewDataStoreReqGetInfo()

	copied.StructureVersion = dataStoreReqGetInfo.StructureVersion

	copied.URL = dataStoreReqGetInfo.URL.Copy().(*types.String)
	copied.RequestHeaders = dataStoreReqGetInfo.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.Size = dataStoreReqGetInfo.Size.Copy().(*types.PrimitiveU32)
	copied.RootCACert = dataStoreReqGetInfo.RootCACert.Copy().(*types.Buffer)
	copied.DataID = dataStoreReqGetInfo.DataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqGetInfo)

	if dataStoreReqGetInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqGetInfo.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqGetInfo.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dataStoreReqGetInfo.Size.Equals(other.Size) {
		return false
	}

	if !dataStoreReqGetInfo.RootCACert.Equals(other.RootCACert) {
		return false
	}

	if !dataStoreReqGetInfo.DataID.Equals(other.DataID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetInfo *DataStoreReqGetInfo) String() string {
	return dataStoreReqGetInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetInfo *DataStoreReqGetInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqGetInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqGetInfo.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dataStoreReqGetInfo.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStoreReqGetInfo.Size))
	b.WriteString(fmt.Sprintf("%sRootCA: %s,\n", indentationValues, dataStoreReqGetInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%sDataID: %s\n", indentationValues, dataStoreReqGetInfo.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	dataStoreReqGetInfo := &DataStoreReqGetInfo{
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		Size:           types.NewPrimitiveU32(0),
		RootCACert:     types.NewBuffer(nil),
		DataID:         types.NewPrimitiveU64(0),
	}

	dataStoreReqGetInfo.RequestHeaders.Type = NewDataStoreKeyValue()

	return dataStoreReqGetInfo
}
