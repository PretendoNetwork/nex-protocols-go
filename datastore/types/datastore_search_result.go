// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSearchResult is a data structure used by the DataStore protocol
type DataStoreSearchResult struct {
	types.Structure
	TotalCount     *types.PrimitiveU32
	Result         *types.List[*DataStoreMetaInfo]
	TotalCountType *types.PrimitiveU8
}

// ExtractFrom extracts the DataStoreSearchResult from the given readable
func (dataStoreSearchResult *DataStoreSearchResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreSearchResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSearchResult header. %s", err.Error())
	}

	err = dataStoreSearchResult.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCount. %s", err.Error())
	}

	err = dataStoreSearchResult.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.Result. %s", err.Error())
	}

	err = dataStoreSearchResult.TotalCountType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCountType. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreSearchResult to the given writable
func (dataStoreSearchResult *DataStoreSearchResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreSearchResult.TotalCount.WriteTo(contentWritable)
	dataStoreSearchResult.Result.WriteTo(contentWritable)
	dataStoreSearchResult.TotalCountType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreSearchResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreSearchResult
func (dataStoreSearchResult *DataStoreSearchResult) Copy() types.RVType {
	copied := NewDataStoreSearchResult()

	copied.StructureVersion = dataStoreSearchResult.StructureVersion

	copied.TotalCount = dataStoreSearchResult.TotalCount.Copy().(*types.PrimitiveU32)
	copied.Result = dataStoreSearchResult.Result.Copy().(*types.List[*DataStoreMetaInfo])
	copied.TotalCountType = dataStoreSearchResult.TotalCountType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchResult *DataStoreSearchResult) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchResult); !ok {
		return false
	}

	other := o.(*DataStoreSearchResult)

	if dataStoreSearchResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreSearchResult.TotalCount.Equals(other.TotalCount) {
		return false
	}

	if !dataStoreSearchResult.Result.Equals(other.Result) {
		return false
	}

	return dataStoreSearchResult.TotalCountType.Equals(other.TotalCountType)
}

// String returns a string representation of the struct
func (dataStoreSearchResult *DataStoreSearchResult) String() string {
	return dataStoreSearchResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchResult *DataStoreSearchResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreSearchResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTotalCount: %s,\n", indentationValues, dataStoreSearchResult.TotalCount))
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, dataStoreSearchResult.Result))
	b.WriteString(fmt.Sprintf("%sTotalCountType: %s\n", indentationValues, dataStoreSearchResult.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchResult returns a new DataStoreSearchResult
func NewDataStoreSearchResult() *DataStoreSearchResult {
	dataStoreSearchResult := &DataStoreSearchResult{
		TotalCount:     types.NewPrimitiveU32(0),
		Result:         types.NewList[*DataStoreMetaInfo](),
		TotalCountType: types.NewPrimitiveU8(0),
	}

	dataStoreSearchResult.Result.Type = NewDataStoreMetaInfo()

	return dataStoreSearchResult
}
