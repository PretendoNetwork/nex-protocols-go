// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetCustomRankingByDataIDParam is a type within the DataStore protocol
type DataStoreGetCustomRankingByDataIDParam struct {
	types.Structure
	ApplicationID types.UInt32
	DataIDList    types.List[types.UInt64]
	ResultOption  types.UInt8
}

// WriteTo writes the DataStoreGetCustomRankingByDataIDParam to the given writable
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgcrbdidp.ApplicationID.WriteTo(contentWritable)
	dsgcrbdidp.DataIDList.WriteTo(contentWritable)
	dsgcrbdidp.ResultOption.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgcrbdidp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetCustomRankingByDataIDParam from the given readable
func (dsgcrbdidp *DataStoreGetCustomRankingByDataIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgcrbdidp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam header. %s", err.Error())
	}

	err = dsgcrbdidp.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ApplicationID. %s", err.Error())
	}

	err = dsgcrbdidp.DataIDList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.DataIDList. %s", err.Error())
	}

	err = dsgcrbdidp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ResultOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCustomRankingByDataIDParam
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) Copy() types.RVType {
	copied := NewDataStoreGetCustomRankingByDataIDParam()

	copied.StructureVersion = dsgcrbdidp.StructureVersion
	copied.ApplicationID = dsgcrbdidp.ApplicationID.Copy().(types.UInt32)
	copied.DataIDList = dsgcrbdidp.DataIDList.Copy().(types.List[types.UInt64])
	copied.ResultOption = dsgcrbdidp.ResultOption.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given DataStoreGetCustomRankingByDataIDParam contains the same data as the current DataStoreGetCustomRankingByDataIDParam
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreGetCustomRankingByDataIDParam); !ok {
		return false
	}

	other := o.(DataStoreGetCustomRankingByDataIDParam)

	if dsgcrbdidp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgcrbdidp.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dsgcrbdidp.DataIDList.Equals(other.DataIDList) {
		return false
	}

	return dsgcrbdidp.ResultOption.Equals(other.ResultOption)
}

// CopyRef copies the current value of the DataStoreGetCustomRankingByDataIDParam
// and returns a pointer to the new copy
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) CopyRef() types.RVTypePtr {
	copied := dsgcrbdidp.Copy().(DataStoreGetCustomRankingByDataIDParam)
	return &copied
}

// Deref takes a pointer to the DataStoreGetCustomRankingByDataIDParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsgcrbdidp *DataStoreGetCustomRankingByDataIDParam) Deref() types.RVType {
	return *dsgcrbdidp
}

// String returns the string representation of the DataStoreGetCustomRankingByDataIDParam
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) String() string {
	return dsgcrbdidp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetCustomRankingByDataIDParam using the provided indentation level
func (dsgcrbdidp DataStoreGetCustomRankingByDataIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingByDataIDParam{\n")
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dsgcrbdidp.ApplicationID))
	b.WriteString(fmt.Sprintf("%sDataIDList: %s,\n", indentationValues, dsgcrbdidp.DataIDList))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgcrbdidp.ResultOption))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingByDataIDParam returns a new DataStoreGetCustomRankingByDataIDParam
func NewDataStoreGetCustomRankingByDataIDParam() DataStoreGetCustomRankingByDataIDParam {
	return DataStoreGetCustomRankingByDataIDParam{
		ApplicationID: types.NewUInt32(0),
		DataIDList:    types.NewList[types.UInt64](),
		ResultOption:  types.NewUInt8(0),
	}

}
