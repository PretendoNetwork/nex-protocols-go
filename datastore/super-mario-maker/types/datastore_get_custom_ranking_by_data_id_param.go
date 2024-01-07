// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreGetCustomRankingByDataIDParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCustomRankingByDataIDParam struct {
	types.Structure
	ApplicationID *types.PrimitiveU32
	DataIDList    *types.List[*types.PrimitiveU64]
	ResultOption  *types.PrimitiveU8
}

// ExtractFrom extracts the DataStoreGetCustomRankingByDataIDParam from the given readable
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreGetCustomRankingByDataIDParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreGetCustomRankingByDataIDParam header. %s", err.Error())
	}

	err = dataStoreGetCustomRankingByDataIDParam.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ApplicationID from stream. %s", err.Error())
	}

	err = dataStoreGetCustomRankingByDataIDParam.DataIDList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.DataIDList from stream. %s", err.Error())
	}

	err = dataStoreGetCustomRankingByDataIDParam.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ResultOption from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreGetCustomRankingByDataIDParam to the given writable
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreGetCustomRankingByDataIDParam.ApplicationID.WriteTo(contentWritable)
	dataStoreGetCustomRankingByDataIDParam.DataIDList.WriteTo(contentWritable)
	dataStoreGetCustomRankingByDataIDParam.ResultOption.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreGetCustomRankingByDataIDParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreGetCustomRankingByDataIDParam
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) Copy() types.RVType {
	copied := NewDataStoreGetCustomRankingByDataIDParam()

	copied.StructureVersion = dataStoreGetCustomRankingByDataIDParam.StructureVersion

	copied.ApplicationID = dataStoreGetCustomRankingByDataIDParam.ApplicationID.Copy().(*types.PrimitiveU32)
	copied.DataIDList = dataStoreGetCustomRankingByDataIDParam.DataIDList.Copy().(*types.List[*types.PrimitiveU64])
	copied.ResultOption = dataStoreGetCustomRankingByDataIDParam.ResultOption.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCustomRankingByDataIDParam); !ok {
		return false
	}

	other := o.(*DataStoreGetCustomRankingByDataIDParam)

	if dataStoreGetCustomRankingByDataIDParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreGetCustomRankingByDataIDParam.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dataStoreGetCustomRankingByDataIDParam.DataIDList.Equals(other.DataIDList) {
		return false
	}

	return dataStoreGetCustomRankingByDataIDParam.ResultOption == other.ResultOption
}

// String returns a string representation of the struct
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) String() string {
	return dataStoreGetCustomRankingByDataIDParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingByDataIDParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.ApplicationID))
	b.WriteString(fmt.Sprintf("%sDataIDList: %s,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.DataIDList))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.ResultOption))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingByDataIDParam returns a new DataStoreGetCustomRankingByDataIDParam
func NewDataStoreGetCustomRankingByDataIDParam() *DataStoreGetCustomRankingByDataIDParam {
	dataStoreGetCustomRankingByDataIDParam := &DataStoreGetCustomRankingByDataIDParam{
		ApplicationID: types.NewPrimitiveU32(0),
		DataIDList:    types.NewList[*types.PrimitiveU64](),
		ResultOption:  types.NewPrimitiveU8(0),
	}

	dataStoreGetCustomRankingByDataIDParam.DataIDList.Type = types.NewPrimitiveU64(0)

	return dataStoreGetCustomRankingByDataIDParam
}
