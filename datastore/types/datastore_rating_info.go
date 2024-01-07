// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingInfo is a data structure used by the DataStore protocol
type DataStoreRatingInfo struct {
	types.Structure
	TotalValue   *types.PrimitiveS64
	Count        *types.PrimitiveU32
	InitialValue *types.PrimitiveS64
}

// ExtractFrom extracts the DataStoreRatingInfo from the given readable
func (dataStoreRatingInfo *DataStoreRatingInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingInfo header. %s", err.Error())
	}

	err = dataStoreRatingInfo.TotalValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.TotalValue. %s", err.Error())
	}

	err = dataStoreRatingInfo.Count.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.Count. %s", err.Error())
	}

	err = dataStoreRatingInfo.InitialValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.InitialValue. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreRatingInfo to the given writable
func (dataStoreRatingInfo *DataStoreRatingInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingInfo.TotalValue.WriteTo(contentWritable)
	dataStoreRatingInfo.Count.WriteTo(contentWritable)
	dataStoreRatingInfo.InitialValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreRatingInfo
func (dataStoreRatingInfo *DataStoreRatingInfo) Copy() types.RVType {
	copied := NewDataStoreRatingInfo()

	copied.StructureVersion = dataStoreRatingInfo.StructureVersion

	copied.TotalValue = dataStoreRatingInfo.TotalValue.Copy().(*types.PrimitiveS64)
	copied.Count = dataStoreRatingInfo.Count.Copy().(*types.PrimitiveU32)
	copied.InitialValue = dataStoreRatingInfo.InitialValue.Copy().(*types.PrimitiveS64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfo *DataStoreRatingInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInfo); !ok {
		return false
	}

	other := o.(*DataStoreRatingInfo)

	if dataStoreRatingInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingInfo.TotalValue.Equals(other.TotalValue) {
		return false
	}

	if !dataStoreRatingInfo.Count.Equals(other.Count) {
		return false
	}

	if !dataStoreRatingInfo.InitialValue.Equals(other.InitialValue) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRatingInfo *DataStoreRatingInfo) String() string {
	return dataStoreRatingInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingInfo *DataStoreRatingInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTotalValue: %s,\n", indentationValues, dataStoreRatingInfo.TotalValue))
	b.WriteString(fmt.Sprintf("%sCount: %s,\n", indentationValues, dataStoreRatingInfo.Count))
	b.WriteString(fmt.Sprintf("%sInitialValue: %s\n", indentationValues, dataStoreRatingInfo.InitialValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfo returns a new DataStoreRatingInfo
func NewDataStoreRatingInfo() *DataStoreRatingInfo {
	return &DataStoreRatingInfo{
		TotalValue:   types.NewPrimitiveS64(0),
		Count:        types.NewPrimitiveU32(0),
		InitialValue: types.NewPrimitiveS64(0),
	}
}
