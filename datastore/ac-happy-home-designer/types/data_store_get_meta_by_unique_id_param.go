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

// WriteTo writes the DataStoreGetMetaByUniqueIdParam to the given variable
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmbuip.UniqueIds.WriteTo(contentWritable)
	dsgmbuip.DataTypes.WriteTo(contentWritable)
	dsgmbuip.ResultOption.WriteTo(contentWritable)
	dsgmbuip.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmbuip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaByUniqueIdParam from the given readable
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmbuip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIdParam header. %s", err.Error())
	}

	err = dsgmbuip.UniqueIds.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIdParam.UniqueIds. %s", err.Error())
	}

	err = dsgmbuip.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIdParam.DataTypes. %s", err.Error())
	}

	err = dsgmbuip.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIdParam.ResultOption. %s", err.Error())
	}

	err = dsgmbuip.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByUniqueIdParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerGetObject
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByUniqueIdParam()

	copied.UniqueIds = dsgmbuip.UniqueIds
	copied.DataTypes = dsgmbuip.DataTypes
	copied.ResultOption = dsgmbuip.ResultOption
	copied.ResultRange = dsgmbuip.ResultRange

	return copied
}

// Equals checks if the given DataStoreGetMetaByUniqueIdParam contains the same data as the current DataStoreGetMetaByUniqueIdParam
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByUniqueIdParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByUniqueIdParam)

	if !dsgmbuip.UniqueIds.Equals(other.UniqueIds) {
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

// String returns the string representation of the DataStoreGetMetaByUniqueIdParam
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) String() string {
	return dsgmbuip.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaByUniqueIdParam using the provided indentation level
func (dsgmbuip *DataStoreGetMetaByUniqueIdParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByUniqueIdParam{\n")
	b.WriteString(fmt.Sprintf("%sUniqueIds: %s,\n", indentationValues, dsgmbuip.UniqueIds))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dsgmbuip.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmbuip.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgmbuip.ResultRange))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByUniqueIdParam returns a new DataStoreGetMetaByUniqueIdParam
func NewDataStoreGetMetaByUniqueIdParam() *DataStoreGetMetaByUniqueIdParam {
	dsgmbuip := &DataStoreGetMetaByUniqueIdParam{
		UniqueIds:    types.NewList[*types.PrimitiveU32](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	return dsgmbuip
}
