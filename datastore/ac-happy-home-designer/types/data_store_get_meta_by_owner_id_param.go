// Package types implements all the types used by the DataStoreACHappyHomeDesigner protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetMetaByOwnerIDParam is a type within the DataStoreACHappyHomeDesigner protocol
type DataStoreGetMetaByOwnerIDParam struct {
	types.Structure
	OwnerIDs     *types.List[*types.PrimitiveU32]
	DataTypes    *types.List[*types.PrimitiveU16]
	ResultOption *types.PrimitiveU8
	ResultRange  *types.ResultRange
}

// WriteTo writes the DataStoreGetMetaByOwnerIDParam to the given variable
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmboip.OwnerIDs.WriteTo(contentWritable)
	dsgmboip.DataTypes.WriteTo(contentWritable)
	dsgmboip.ResultOption.WriteTo(contentWritable)
	dsgmboip.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmboip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaByOwnerIDParam from the given readable
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmboip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam header. %s", err.Error())
	}

	err = dsgmboip.OwnerIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.OwnerIDs. %s", err.Error())
	}

	err = dsgmboip.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.DataTypes. %s", err.Error())
	}

	err = dsgmboip.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption. %s", err.Error())
	}

	err = dsgmboip.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerGetObject
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByOwnerIDParam()

	copied.OwnerIDs = dsgmboip.OwnerIDs
	copied.DataTypes = dsgmboip.DataTypes
	copied.ResultOption = dsgmboip.ResultOption
	copied.ResultRange = dsgmboip.ResultRange

	return copied
}

// Equals checks if the given DataStoreGetMetaByOwnerIDParam contains the same data as the current DataStoreGetMetaByOwnerIDParam
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByOwnerIDParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByOwnerIDParam)

	if !dsgmboip.OwnerIDs.Equals(other.OwnerIDs) {
		return false
	}

	if !dsgmboip.DataTypes.Equals(other.DataTypes) {
		return false
	}

	if !dsgmboip.ResultOption.Equals(other.ResultOption) {
		return false
	}

	return dsgmboip.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the DataStoreGetMetaByOwnerIDParam
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) String() string {
	return dsgmboip.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaByOwnerIDParam using the provided indentation level
func (dsgmboip *DataStoreGetMetaByOwnerIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByOwnerIDParam{\n")
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %s,\n", indentationValues, dsgmboip.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dsgmboip.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmboip.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgmboip.ResultRange))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	dsgmboip := &DataStoreGetMetaByOwnerIDParam{
		OwnerIDs:     types.NewList[*types.PrimitiveU32](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	return dsgmboip
}
