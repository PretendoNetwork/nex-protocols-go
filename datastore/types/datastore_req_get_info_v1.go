// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReqGetInfoV1 is a data structure used by the DataStore protocol
type DataStoreReqGetInfoV1 struct {
	types.Structure
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	Size           *types.PrimitiveU32
	RootCACert     *types.Buffer
}

// ExtractFrom extracts the DataStoreReqGetInfoV1 from the given readable
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReqGetInfoV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReqGetInfoV1 header. %s", err.Error())
	}

	err = dataStoreReqGetInfoV1.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.URL. %s", err.Error())
	}

	err = dataStoreReqGetInfoV1.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RequestHeaders. %s", err.Error())
	}

	err = dataStoreReqGetInfoV1.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.Size. %s", err.Error())
	}

	err = dataStoreReqGetInfoV1.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RootCACert. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReqGetInfoV1 to the given writable
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReqGetInfoV1.URL.WriteTo(contentWritable)
	dataStoreReqGetInfoV1.RequestHeaders.WriteTo(contentWritable)
	dataStoreReqGetInfoV1.Size.WriteTo(contentWritable)
	dataStoreReqGetInfoV1.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReqGetInfoV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReqGetInfoV1
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Copy() types.RVType {
	copied := NewDataStoreReqGetInfoV1()

	copied.StructureVersion = dataStoreReqGetInfoV1.StructureVersion

	copied.URL = dataStoreReqGetInfoV1.URL.Copy().(*types.String)
	copied.RequestHeaders = dataStoreReqGetInfoV1.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.Size = dataStoreReqGetInfoV1.Size.Copy().(*types.PrimitiveU32)
	copied.RootCACert = dataStoreReqGetInfoV1.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetInfoV1); !ok {
		return false
	}

	other := o.(*DataStoreReqGetInfoV1)

	if dataStoreReqGetInfoV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReqGetInfoV1.URL.Equals(other.URL) {
		return false
	}

	if !dataStoreReqGetInfoV1.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dataStoreReqGetInfoV1.Size.Equals(other.Size) {
		return false
	}

	if !dataStoreReqGetInfoV1.RootCACert.Equals(other.RootCACert) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) String() string {
	return dataStoreReqGetInfoV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReqGetInfoV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dataStoreReqGetInfoV1.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dataStoreReqGetInfoV1.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStoreReqGetInfoV1.Size))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s\n", indentationValues, dataStoreReqGetInfoV1.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfoV1 returns a new DataStoreReqGetInfoV1
func NewDataStoreReqGetInfoV1() *DataStoreReqGetInfoV1 {
	dataStoreReqGetInfoV1 := &DataStoreReqGetInfoV1{
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		Size:           types.NewPrimitiveU32(0),
		RootCACert:     types.NewBuffer(nil),
	}

	dataStoreReqGetInfoV1.RequestHeaders.Type = NewDataStoreKeyValue()

	return dataStoreReqGetInfoV1
}
