// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreChangePlayablePlatformParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreChangePlayablePlatformParam struct {
	types.Structure
	DataID           *types.PrimitiveU64
	PlayablePlatform *types.PrimitiveU32
}

// ExtractFrom extracts the DataStoreChangePlayablePlatformParam from the given readable
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreChangePlayablePlatformParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreChangePlayablePlatformParam header. %s", err.Error())
	}

	err = dataStoreChangePlayablePlatformParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.DataID from stream. %s", err.Error())
	}

	err = dataStoreChangePlayablePlatformParam.PlayablePlatform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.PlayablePlatform from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreChangePlayablePlatformParam to the given writable
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreChangePlayablePlatformParam.DataID.WriteTo(contentWritable)
	dataStoreChangePlayablePlatformParam.PlayablePlatform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreChangePlayablePlatformParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreChangePlayablePlatformParam
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) Copy() types.RVType {
	copied := NewDataStoreChangePlayablePlatformParam()

	copied.StructureVersion = dataStoreChangePlayablePlatformParam.StructureVersion

	copied.DataID = dataStoreChangePlayablePlatformParam.DataID.Copy().(*types.PrimitiveU64)
	copied.PlayablePlatform = dataStoreChangePlayablePlatformParam.PlayablePlatform.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangePlayablePlatformParam); !ok {
		return false
	}

	other := o.(*DataStoreChangePlayablePlatformParam)

	if dataStoreChangePlayablePlatformParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreChangePlayablePlatformParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreChangePlayablePlatformParam.PlayablePlatform.Equals(other.PlayablePlatform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) String() string {
	return dataStoreChangePlayablePlatformParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangePlayablePlatformParam *DataStoreChangePlayablePlatformParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangePlayablePlatformParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreChangePlayablePlatformParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreChangePlayablePlatformParam.DataID))
	b.WriteString(fmt.Sprintf("%sPlayablePlatform: %s,\n", indentationValues, dataStoreChangePlayablePlatformParam.PlayablePlatform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangePlayablePlatformParam returns a new DataStoreChangePlayablePlatformParam
func NewDataStoreChangePlayablePlatformParam() *DataStoreChangePlayablePlatformParam {
	return &DataStoreChangePlayablePlatformParam{
		DataID:           types.NewPrimitiveU64(0),
		PlayablePlatform: types.NewPrimitiveU32(0),
	}
}
