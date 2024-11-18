// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetCustomRankingParam is a type within the DataStore protocol
type DataStoreGetCustomRankingParam struct {
	types.Structure
	ApplicationID types.UInt32
	Condition     DataStoreCustomRankingRatingCondition
	ResultOption  types.UInt8
	ResultRange   types.ResultRange
}

// WriteTo writes the DataStoreGetCustomRankingParam to the given writable
func (dsgcrp DataStoreGetCustomRankingParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgcrp.ApplicationID.WriteTo(contentWritable)
	dsgcrp.Condition.WriteTo(contentWritable)
	dsgcrp.ResultOption.WriteTo(contentWritable)
	dsgcrp.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgcrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetCustomRankingParam from the given readable
func (dsgcrp *DataStoreGetCustomRankingParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgcrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam header. %s", err.Error())
	}

	err = dsgcrp.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ApplicationID. %s", err.Error())
	}

	err = dsgcrp.Condition.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.Condition. %s", err.Error())
	}

	err = dsgcrp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultOption. %s", err.Error())
	}

	err = dsgcrp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCustomRankingParam
func (dsgcrp DataStoreGetCustomRankingParam) Copy() types.RVType {
	copied := NewDataStoreGetCustomRankingParam()

	copied.StructureVersion = dsgcrp.StructureVersion
	copied.ApplicationID = dsgcrp.ApplicationID.Copy().(types.UInt32)
	copied.Condition = dsgcrp.Condition.Copy().(DataStoreCustomRankingRatingCondition)
	copied.ResultOption = dsgcrp.ResultOption.Copy().(types.UInt8)
	copied.ResultRange = dsgcrp.ResultRange.Copy().(types.ResultRange)

	return copied
}

// Equals checks if the given DataStoreGetCustomRankingParam contains the same data as the current DataStoreGetCustomRankingParam
func (dsgcrp DataStoreGetCustomRankingParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCustomRankingParam); !ok {
		return false
	}

	other := o.(*DataStoreGetCustomRankingParam)

	if dsgcrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgcrp.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dsgcrp.Condition.Equals(other.Condition) {
		return false
	}

	if !dsgcrp.ResultOption.Equals(other.ResultOption) {
		return false
	}

	return dsgcrp.ResultRange.Equals(other.ResultRange)
}

// CopyRef copies the current value of the DataStoreGetCustomRankingParam
// and returns a pointer to the new copy
func (dsgcrp DataStoreGetCustomRankingParam) CopyRef() types.RVTypePtr {
	copied := dsgcrp.Copy().(DataStoreGetCustomRankingParam)
	return &copied
}

// Deref takes a pointer to the DataStoreGetCustomRankingParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsgcrp *DataStoreGetCustomRankingParam) Deref() types.RVType {
	return *dsgcrp
}

// String returns the string representation of the DataStoreGetCustomRankingParam
func (dsgcrp DataStoreGetCustomRankingParam) String() string {
	return dsgcrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetCustomRankingParam using the provided indentation level
func (dsgcrp DataStoreGetCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dsgcrp.ApplicationID))
	b.WriteString(fmt.Sprintf("%sCondition: %s,\n", indentationValues, dsgcrp.Condition.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgcrp.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgcrp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingParam returns a new DataStoreGetCustomRankingParam
func NewDataStoreGetCustomRankingParam() DataStoreGetCustomRankingParam {
	return DataStoreGetCustomRankingParam{
		ApplicationID: types.NewUInt32(0),
		Condition:     NewDataStoreCustomRankingRatingCondition(),
		ResultOption:  types.NewUInt8(0),
		ResultRange:   types.NewResultRange(),
	}

}
