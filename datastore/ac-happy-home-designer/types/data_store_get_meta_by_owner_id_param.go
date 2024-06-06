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

// WriteTo writes the DataStoreGetMetaByOwnerIdParam to the given variable
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmboip.OwnerIds.WriteTo(contentWritable)
	dsgmboip.DataTypes.WriteTo(contentWritable)
	dsgmboip.ResultOption.WriteTo(contentWritable)
	dsgmboip.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmboip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaByOwnerIdParam from the given readable
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmboip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIdParam header. %s", err.Error())
	}

	err = dsgmboip.OwnerIds.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIdParam.OwnerIds. %s", err.Error())
	}

	err = dsgmboip.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIdParam.DataTypes. %s", err.Error())
	}

	err = dsgmboip.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIdParam.ResultOption. %s", err.Error())
	}

	err = dsgmboip.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIdParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerGetObject
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByOwnerIdParam()

	copied.OwnerIds = dsgmboip.OwnerIds
	copied.DataTypes = dsgmboip.DataTypes
	copied.ResultOption = dsgmboip.ResultOption
	copied.ResultRange = dsgmboip.ResultRange

	return copied
}

// Equals checks if the given DataStoreGetMetaByOwnerIdParam contains the same data as the current DataStoreGetMetaByOwnerIdParam
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByOwnerIdParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByOwnerIdParam)

	if !dsgmboip.OwnerIds.Equals(other.OwnerIds) {
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

// String returns the string representation of the DataStoreGetMetaByOwnerIdParam
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) String() string {
	return dsgmboip.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaByOwnerIdParam using the provided indentation level
func (dsgmboip *DataStoreGetMetaByOwnerIdParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByOwnerIdParam{\n")
	b.WriteString(fmt.Sprintf("%sOwnerIds: %s,\n", indentationValues, dsgmboip.OwnerIds))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dsgmboip.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmboip.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgmboip.ResultRange))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByOwnerIdParam returns a new DataStoreGetMetaByOwnerIdParam
func NewDataStoreGetMetaByOwnerIdParam() *DataStoreGetMetaByOwnerIdParam {
	dsgmboip := &DataStoreGetMetaByOwnerIdParam{
		OwnerIds:     types.NewList[*types.PrimitiveU32](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	return dsgmboip
}
