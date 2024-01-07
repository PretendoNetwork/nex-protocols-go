// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	types.Structure
	SearchTarget           *types.PrimitiveU8
	OwnerIDs               *types.List[*types.PID]
	OwnerType              *types.PrimitiveU8
	DestinationIDs         *types.List[*types.PID]
	DataType               *types.PrimitiveU16
	CreatedAfter           *types.DateTime
	CreatedBefore          *types.DateTime
	UpdatedAfter           *types.DateTime
	UpdatedBefore          *types.DateTime
	ReferDataID            *types.PrimitiveU32
	Tags                   *types.List[*types.String]
	ResultOrderColumn      *types.PrimitiveU8
	ResultOrder            *types.PrimitiveU8
	ResultRange            *types.ResultRange
	ResultOption           *types.PrimitiveU8
	MinimalRatingFrequency *types.PrimitiveU32
	UseCache               *types.PrimitiveBool
	TotalCountEnabled      *types.PrimitiveBool
	DataTypes              *types.List[*types.PrimitiveU16]
}

// WriteTo writes the DataStoreSearchParam to the given writable
func (dataStoreSearchParam *DataStoreSearchParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreSearchParam.SearchTarget.WriteTo(contentWritable)
	dataStoreSearchParam.OwnerIDs.WriteTo(contentWritable)
	dataStoreSearchParam.OwnerType.WriteTo(contentWritable)
	dataStoreSearchParam.DestinationIDs.WriteTo(contentWritable)
	dataStoreSearchParam.DataType.WriteTo(contentWritable)
	dataStoreSearchParam.CreatedAfter.WriteTo(contentWritable)
	dataStoreSearchParam.CreatedBefore.WriteTo(contentWritable)
	dataStoreSearchParam.UpdatedAfter.WriteTo(contentWritable)
	dataStoreSearchParam.UpdatedBefore.WriteTo(contentWritable)
	dataStoreSearchParam.ReferDataID.WriteTo(contentWritable)
	dataStoreSearchParam.Tags.WriteTo(contentWritable)
	dataStoreSearchParam.ResultOrderColumn.WriteTo(contentWritable)
	dataStoreSearchParam.ResultOrder.WriteTo(contentWritable)
	dataStoreSearchParam.ResultRange.WriteTo(contentWritable)
	dataStoreSearchParam.ResultOption.WriteTo(contentWritable)
	dataStoreSearchParam.MinimalRatingFrequency.WriteTo(contentWritable)

	if dataStoreSearchParam.StructureVersion >= 1 {
		dataStoreSearchParam.UseCache.WriteTo(contentWritable)
	}

	if dataStoreSearchParam.StructureVersion >= 3 {
		dataStoreSearchParam.TotalCountEnabled.WriteTo(contentWritable)
	}

	if dataStoreSearchParam.StructureVersion >= 2 {
		dataStoreSearchParam.DataTypes.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStoreSearchParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchParam from the given readable
func (dataStoreSearchParam *DataStoreSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreSearchParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSearchParam header. %s", err.Error())
	}

	err = dataStoreSearchParam.SearchTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.SearchTarget. %s", err.Error())
	}

	err = dataStoreSearchParam.OwnerIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerIDs. %s", err.Error())
	}

	err = dataStoreSearchParam.OwnerType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerType. %s", err.Error())
	}

	err = dataStoreSearchParam.DestinationIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DestinationIDs. %s", err.Error())
	}

	err = dataStoreSearchParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DataType. %s", err.Error())
	}

	err = dataStoreSearchParam.CreatedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedAfter. %s", err.Error())
	}

	err = dataStoreSearchParam.CreatedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedBefore. %s", err.Error())
	}

	err = dataStoreSearchParam.UpdatedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedAfter. %s", err.Error())
	}

	err = dataStoreSearchParam.UpdatedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedBefore. %s", err.Error())
	}

	err = dataStoreSearchParam.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ReferDataID. %s", err.Error())
	}

	err = dataStoreSearchParam.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.Tags. %s", err.Error())
	}

	err = dataStoreSearchParam.ResultOrderColumn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrderColumn. %s", err.Error())
	}

	err = dataStoreSearchParam.ResultOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrder. %s", err.Error())
	}

	err = dataStoreSearchParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultRange. %s", err.Error())
	}

	err = dataStoreSearchParam.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOption. %s", err.Error())
	}

	err = dataStoreSearchParam.MinimalRatingFrequency.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.MinimalRatingFrequency. %s", err.Error())
	}

	if dataStoreSearchParam.StructureVersion >= 1 {
		err = dataStoreSearchParam.UseCache.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.UseCache. %s", err.Error())
		}
	}

	if dataStoreSearchParam.StructureVersion >= 3 {
		err = dataStoreSearchParam.TotalCountEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.TotalCountEnabled. %s", err.Error())
		}
	}

	if dataStoreSearchParam.StructureVersion >= 2 {
		err = dataStoreSearchParam.DataTypes.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.DataTypes. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchParam
func (dataStoreSearchParam *DataStoreSearchParam) Copy() types.RVType {
	copied := NewDataStoreSearchParam()

	copied.StructureVersion = dataStoreSearchParam.StructureVersion

	copied.SearchTarget = dataStoreSearchParam.SearchTarget.Copy().(*types.PrimitiveU8)
	copied.OwnerIDs = dataStoreSearchParam.OwnerIDs.Copy().(*types.List[*types.PID])
	copied.OwnerType = dataStoreSearchParam.OwnerType.Copy().(*types.PrimitiveU8)
	copied.DestinationIDs = dataStoreSearchParam.DestinationIDs.Copy().(*types.List[*types.PID])
	copied.DataType = dataStoreSearchParam.DataType.Copy().(*types.PrimitiveU16)
	copied.CreatedAfter = dataStoreSearchParam.CreatedAfter.Copy().(*types.DateTime)
	copied.CreatedBefore = dataStoreSearchParam.CreatedBefore.Copy().(*types.DateTime)
	copied.UpdatedAfter = dataStoreSearchParam.UpdatedAfter.Copy().(*types.DateTime)
	copied.UpdatedBefore = dataStoreSearchParam.UpdatedBefore.Copy().(*types.DateTime)
	copied.ReferDataID = dataStoreSearchParam.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dataStoreSearchParam.Tags.Copy().(*types.List[*types.String])
	copied.ResultOrderColumn = dataStoreSearchParam.ResultOrderColumn.Copy().(*types.PrimitiveU8)
	copied.ResultOrder = dataStoreSearchParam.ResultOrder.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dataStoreSearchParam.ResultRange.Copy().(*types.ResultRange)
	copied.ResultOption = dataStoreSearchParam.ResultOption.Copy().(*types.PrimitiveU8)
	copied.MinimalRatingFrequency = dataStoreSearchParam.MinimalRatingFrequency.Copy().(*types.PrimitiveU32)
	copied.UseCache = dataStoreSearchParam.UseCache.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchParam *DataStoreSearchParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchParam)

	if dataStoreSearchParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreSearchParam.SearchTarget.Equals(other.SearchTarget) {
		return false
	}

	if !dataStoreSearchParam.OwnerIDs.Equals(other.OwnerIDs) {
		return false
	}

	if !dataStoreSearchParam.OwnerType.Equals(other.OwnerType) {
		return false
	}

	if !dataStoreSearchParam.DestinationIDs.Equals(other.DestinationIDs) {
		return false
	}

	if !dataStoreSearchParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreSearchParam.CreatedAfter.Equals(other.CreatedAfter) {
		return false
	}

	if !dataStoreSearchParam.CreatedBefore.Equals(other.CreatedBefore) {
		return false
	}

	if !dataStoreSearchParam.UpdatedAfter.Equals(other.UpdatedAfter) {
		return false
	}

	if !dataStoreSearchParam.UpdatedBefore.Equals(other.UpdatedBefore) {
		return false
	}

	if !dataStoreSearchParam.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dataStoreSearchParam.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStoreSearchParam.ResultOrderColumn.Equals(other.ResultOrderColumn) {
		return false
	}

	if !dataStoreSearchParam.ResultOrder.Equals(other.ResultOrder) {
		return false
	}

	if !dataStoreSearchParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if !dataStoreSearchParam.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dataStoreSearchParam.MinimalRatingFrequency.Equals(other.MinimalRatingFrequency) {
		return false
	}

	if !dataStoreSearchParam.UseCache.Equals(other.UseCache) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSearchParam *DataStoreSearchParam) String() string {
	return dataStoreSearchParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchParam *DataStoreSearchParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreSearchParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSearchTarget: %s,\n", indentationValues, dataStoreSearchParam.SearchTarget))
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %s,\n", indentationValues, dataStoreSearchParam.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sOwnerType: %s,\n", indentationValues, dataStoreSearchParam.OwnerType))
	b.WriteString(fmt.Sprintf("%sDestinationIDs: %s,\n", indentationValues, dataStoreSearchParam.DestinationIDs))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStoreSearchParam.DataType))
	b.WriteString(fmt.Sprintf("%sCreatedAfter: %s,\n", indentationValues, dataStoreSearchParam.CreatedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCreatedBefore: %s,\n", indentationValues, dataStoreSearchParam.CreatedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedAfter: %s,\n", indentationValues, dataStoreSearchParam.UpdatedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedBefore: %s,\n", indentationValues, dataStoreSearchParam.UpdatedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dataStoreSearchParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStoreSearchParam.Tags))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %s,\n", indentationValues, dataStoreSearchParam.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %s,\n", indentationValues, dataStoreSearchParam.ResultOrder))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dataStoreSearchParam.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dataStoreSearchParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sMinimalRatingFrequency: %s,\n", indentationValues, dataStoreSearchParam.MinimalRatingFrequency))
	b.WriteString(fmt.Sprintf("%sUseCache: %s,\n", indentationValues, dataStoreSearchParam.UseCache))
	b.WriteString(fmt.Sprintf("%sTotalCountEnabled: %s,\n", indentationValues, dataStoreSearchParam.TotalCountEnabled))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s\n", indentationValues, dataStoreSearchParam.DataTypes))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	dataStoreSearchParam := &DataStoreSearchParam{
		SearchTarget:           types.NewPrimitiveU8(0),
		OwnerIDs:               types.NewList[*types.PID](),
		OwnerType:              types.NewPrimitiveU8(0),
		DestinationIDs:         types.NewList[*types.PID](),
		DataType:               types.NewPrimitiveU16(0),
		CreatedAfter:           types.NewDateTime(0),
		CreatedBefore:          types.NewDateTime(0),
		UpdatedAfter:           types.NewDateTime(0),
		UpdatedBefore:          types.NewDateTime(0),
		ReferDataID:            types.NewPrimitiveU32(0),
		Tags:                   types.NewList[*types.String](),
		ResultOrderColumn:      types.NewPrimitiveU8(0),
		ResultOrder:            types.NewPrimitiveU8(0),
		ResultRange:            types.NewResultRange(),
		ResultOption:           types.NewPrimitiveU8(0),
		MinimalRatingFrequency: types.NewPrimitiveU32(0),
		UseCache:               types.NewPrimitiveBool(false),
		TotalCountEnabled:      types.NewPrimitiveBool(false),
		DataTypes:              types.NewList[*types.PrimitiveU16](),
	}

	dataStoreSearchParam.OwnerIDs.Type = types.NewPID(0)
	dataStoreSearchParam.DestinationIDs.Type = types.NewPID(0)
	dataStoreSearchParam.Tags.Type = types.NewString("")
	dataStoreSearchParam.DataTypes.Type = types.NewPrimitiveU16(0)

	return dataStoreSearchParam
}
