// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetCustomRankingParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCustomRankingParam struct {
	types.Structure
	ApplicationID *types.PrimitiveU32
	Condition     *DataStoreCustomRankingRatingCondition
	ResultOption  *types.PrimitiveU8
	ResultRange   *types.ResultRange
}

// ExtractFrom extracts the DataStoreGetCustomRankingParam from the given readable
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetCustomRankingParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetCustomRankingParam header. %s", err.Error())
	}

	err = dataStoreGetCustomRankingParam.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ApplicationID from stream. %s", err.Error())
	}

	err = dataStoreGetCustomRankingParam.Condition.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.Condition from stream. %s", err.Error())
	}

	err = dataStoreGetCustomRankingParam.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultOption from stream. %s", err.Error())
	}

	err = dataStoreGetCustomRankingParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultRange from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetCustomRankingParam to the given writable
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetCustomRankingParam.ApplicationID.WriteTo(contentWritable)
	dataStoreGetCustomRankingParam.Condition.WriteTo(contentWritable)
	dataStoreGetCustomRankingParam.ResultOption.WriteTo(contentWritable)
	dataStoreGetCustomRankingParam.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetCustomRankingParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetCustomRankingParam
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) Copy() types.RVType {
	copied := NewDataStoreGetCustomRankingParam()

	copied.StructureVersion = dataStoreGetCustomRankingParam.StructureVersion

	copied.ApplicationID = dataStoreGetCustomRankingParam.ApplicationID.Copy().(*types.PrimitiveU32)
	copied.Condition = dataStoreGetCustomRankingParam.Condition.Copy().(*DataStoreCustomRankingRatingCondition)
	copied.ResultOption = dataStoreGetCustomRankingParam.ResultOption.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dataStoreGetCustomRankingParam.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCustomRankingParam); !ok {
		return false
	}

	other := o.(*DataStoreGetCustomRankingParam)

	if dataStoreGetCustomRankingParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetCustomRankingParam.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dataStoreGetCustomRankingParam.Condition.Equals(other.Condition) {
		return false
	}

	if !dataStoreGetCustomRankingParam.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dataStoreGetCustomRankingParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) String() string {
	return dataStoreGetCustomRankingParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetCustomRankingParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dataStoreGetCustomRankingParam.ApplicationID))
	b.WriteString(fmt.Sprintf("%sCondition: %s\n", indentationValues, dataStoreGetCustomRankingParam.Condition.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dataStoreGetCustomRankingParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreGetCustomRankingParam.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingParam returns a new DataStoreGetCustomRankingParam
func NewDataStoreGetCustomRankingParam() *DataStoreGetCustomRankingParam {
	return &DataStoreGetCustomRankingParam{
		ApplicationID: types.NewPrimitiveU32(0),
		Condition:     NewDataStoreCustomRankingRatingCondition(),
		ResultOption:  types.NewPrimitiveU8(0),
		ResultRange:   types.NewResultRange(),
	}
}
