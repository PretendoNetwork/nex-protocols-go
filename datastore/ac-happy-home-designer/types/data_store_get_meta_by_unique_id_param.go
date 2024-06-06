// Package types implements all the types used by the DataStoreACHappyHomeDesigner protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetMetaByUniqueIDParam is a type within the DataStoreACHappyHomeDesigner protocol
type DataStoreGetMetaByUniqueIDParam struct {
	types.Structure
	UniqueIDs    *types.List[*types.PrimitiveU32]
	DataTypes    *types.List[*types.PrimitiveU16]
	ResultOption *types.PrimitiveU8
	ResultRange  *types.ResultRange
}

// WriteTo writes the DataStoreGetMetaByUniqueIDParam to the given variable
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmbuip.UniqueIDs.WriteTo(contentWritable)
	dsgmbuip.DataTypes.WriteTo(contentWritable)
	dsgmbuip.ResultOption.WriteTo(contentWritable)
	dsgmbuip.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmbuip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaByUniqueIDParam from the given readable
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmbuip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIDParam header. %s", err.Error())
	}

	err = dsgmbuip.UniqueIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIDParam.UniqueIDs. %s", err.Error())
	}

	err = dsgmbuip.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIDParam.DataTypes. %s", err.Error())
	}

	err = dsgmbuip.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIDParam.ResultOption. %s", err.Error())
	}

	err = dsgmbuip.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIDParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerGetObject
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByUniqueIDParam()

	copied.UniqueIDs = dsgmbuip.UniqueIDs
	copied.DataTypes = dsgmbuip.DataTypes
	copied.ResultOption = dsgmbuip.ResultOption
	copied.ResultRange = dsgmbuip.ResultRange

	return copied
}

// Equals checks if the given DataStoreGetMetaByUniqueIDParam contains the same data as the current DataStoreGetMetaByUniqueIDParam
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByUniqueIDParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByUniqueIDParam)

	if !dsgmbuip.UniqueIDs.Equals(other.UniqueIDs) {
		return false
	}

	if !dsgmbuip.DataTypes.Equals(other.DataTypes) {
		return false
	}

	if !dsgmbuip.ResultOption.Equals(other.ResultOption) {
		return false
	}

	return dsgmbuip.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the DataStoreGetMetaByUniqueIDParam
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) String() string {
	return dsgmbuip.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaByUniqueIDParam using the provided indentation level
func (dsgmbuip *DataStoreGetMetaByUniqueIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByUniqueIDParam{\n")
	b.WriteString(fmt.Sprintf("%sUniqueIDs: %s,\n", indentationValues, dsgmbuip.UniqueIDs))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dsgmbuip.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmbuip.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgmbuip.ResultRange))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByUniqueIDParam returns a new DataStoreGetMetaByUniqueIDParam
func NewDataStoreGetMetaByUniqueIDParam() *DataStoreGetMetaByUniqueIDParam {
	dsgmbuip := &DataStoreGetMetaByUniqueIDParam{
		UniqueIDs:    types.NewList[*types.PrimitiveU32](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	return dsgmbuip
}
