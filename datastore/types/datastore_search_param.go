// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	nex.Structure
	SearchTarget           uint8
	OwnerIDs               []uint32
	OwnerType              uint8
	DestinationIDs         []uint64
	DataType               uint16
	CreatedAfter           *nex.DateTime
	CreatedBefore          *nex.DateTime
	UpdatedAfter           *nex.DateTime
	UpdatedBefore          *nex.DateTime
	ReferDataID            uint32
	Tags                   []string
	ResultOrderColumn      uint8
	ResultOrder            uint8
	ResultRange            *nex.ResultRange
	ResultOption           uint8
	MinimalRatingFrequency uint32
	UseCache               bool
	TotalCountEnabled      bool
	DataTypes              []uint16
}

// ExtractFromStream extracts a DataStoreSearchParam structure from a stream
func (dataStoreSearchParam *DataStoreSearchParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSearchParam.SearchTarget, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.SearchTarget. %s", err.Error())
	}

	dataStoreSearchParam.OwnerIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerIDs. %s", err.Error())
	}

	dataStoreSearchParam.OwnerType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerType. %s", err.Error())
	}

	dataStoreSearchParam.DestinationIDs, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DestinationIDs. %s", err.Error())
	}

	dataStoreSearchParam.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DataType. %s", err.Error())
	}

	dataStoreSearchParam.CreatedAfter, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedAfter. %s", err.Error())
	}

	dataStoreSearchParam.CreatedBefore, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedBefore. %s", err.Error())
	}

	dataStoreSearchParam.UpdatedAfter, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedAfter. %s", err.Error())
	}

	dataStoreSearchParam.UpdatedBefore, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedBefore. %s", err.Error())
	}

	dataStoreSearchParam.ReferDataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ReferDataID. %s", err.Error())
	}

	dataStoreSearchParam.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.Tags. %s", err.Error())
	}

	dataStoreSearchParam.ResultOrderColumn, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrderColumn. %s", err.Error())
	}

	dataStoreSearchParam.ResultOrder, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrder. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultRange. %s", err.Error())
	}

	dataStoreSearchParam.ResultRange = resultRange.(*nex.ResultRange)
	dataStoreSearchParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOption. %s", err.Error())
	}

	dataStoreSearchParam.MinimalRatingFrequency, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.MinimalRatingFrequency. %s", err.Error())
	}

	if dataStoreSearchParam.StructureVersion() >= 1 {
		dataStoreSearchParam.UseCache, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.UseCache. %s", err.Error())
		}
	}

	if dataStoreSearchParam.StructureVersion() >= 3 {
		dataStoreSearchParam.TotalCountEnabled, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.TotalCountEnabled. %s", err.Error())
		}
	}

	if dataStoreSearchParam.StructureVersion() >= 2 {
		dataStoreSearchParam.DataTypes, err = stream.ReadListUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreSearchParam.DataTypes. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchParam
func (dataStoreSearchParam *DataStoreSearchParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchParam()

	copied.SearchTarget = dataStoreSearchParam.SearchTarget
	copied.OwnerIDs = make([]uint32, len(dataStoreSearchParam.OwnerIDs))

	copy(copied.OwnerIDs, dataStoreSearchParam.OwnerIDs)

	copied.OwnerType = dataStoreSearchParam.OwnerType
	copied.DestinationIDs = make([]uint64, len(dataStoreSearchParam.DestinationIDs))

	copy(copied.DestinationIDs, dataStoreSearchParam.DestinationIDs)

	copied.DataType = dataStoreSearchParam.DataType
	copied.CreatedAfter = dataStoreSearchParam.CreatedAfter.Copy()
	copied.CreatedBefore = dataStoreSearchParam.CreatedBefore.Copy()
	copied.UpdatedAfter = dataStoreSearchParam.UpdatedAfter.Copy()
	copied.UpdatedBefore = dataStoreSearchParam.UpdatedBefore.Copy()
	copied.ReferDataID = dataStoreSearchParam.ReferDataID
	copied.Tags = make([]string, len(dataStoreSearchParam.Tags))

	copy(copied.Tags, dataStoreSearchParam.Tags)

	copied.ResultOrderColumn = dataStoreSearchParam.ResultOrderColumn
	copied.ResultOrder = dataStoreSearchParam.ResultOrder
	copied.ResultRange = dataStoreSearchParam.ResultRange.Copy().(*nex.ResultRange)
	copied.ResultOption = dataStoreSearchParam.ResultOption
	copied.MinimalRatingFrequency = dataStoreSearchParam.MinimalRatingFrequency
	copied.UseCache = dataStoreSearchParam.UseCache

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchParam *DataStoreSearchParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchParam)

	if dataStoreSearchParam.SearchTarget != other.SearchTarget {
		return false
	}

	if len(dataStoreSearchParam.OwnerIDs) != len(other.OwnerIDs) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.OwnerIDs); i++ {
		if dataStoreSearchParam.OwnerIDs[i] != other.OwnerIDs[i] {
			return false
		}
	}

	if dataStoreSearchParam.OwnerType != other.OwnerType {
		return false
	}

	if len(dataStoreSearchParam.DestinationIDs) != len(other.DestinationIDs) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.DestinationIDs); i++ {
		if dataStoreSearchParam.DestinationIDs[i] != other.DestinationIDs[i] {
			return false
		}
	}

	if dataStoreSearchParam.DataType != other.DataType {
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

	if dataStoreSearchParam.ReferDataID != other.ReferDataID {
		return false
	}

	if len(dataStoreSearchParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.Tags); i++ {
		if dataStoreSearchParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreSearchParam.ResultOrderColumn != other.ResultOrderColumn {
		return false
	}

	if dataStoreSearchParam.ResultOrder != other.ResultOrder {
		return false
	}

	if !dataStoreSearchParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if dataStoreSearchParam.ResultOption != other.ResultOption {
		return false
	}

	if dataStoreSearchParam.MinimalRatingFrequency != other.MinimalRatingFrequency {
		return false
	}

	if dataStoreSearchParam.UseCache != other.UseCache {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSearchParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSearchTarget: %d,\n", indentationValues, dataStoreSearchParam.SearchTarget))
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %v,\n", indentationValues, dataStoreSearchParam.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sOwnerType: %d,\n", indentationValues, dataStoreSearchParam.OwnerType))
	b.WriteString(fmt.Sprintf("%sDestinationIDs: %v,\n", indentationValues, dataStoreSearchParam.DestinationIDs))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSearchParam.DataType))

	if dataStoreSearchParam.CreatedAfter != nil {
		b.WriteString(fmt.Sprintf("%sCreatedAfter: %s,\n", indentationValues, dataStoreSearchParam.CreatedAfter.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedAfter: nil,\n", indentationValues))
	}

	if dataStoreSearchParam.CreatedBefore != nil {
		b.WriteString(fmt.Sprintf("%sCreatedBefore: %s,\n", indentationValues, dataStoreSearchParam.CreatedBefore.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedBefore: nil,\n", indentationValues))
	}

	if dataStoreSearchParam.UpdatedAfter != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedAfter: %s,\n", indentationValues, dataStoreSearchParam.UpdatedAfter.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedAfter: nil,\n", indentationValues))
	}

	if dataStoreSearchParam.UpdatedBefore != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedBefore: %s,\n", indentationValues, dataStoreSearchParam.UpdatedBefore.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedBefore: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sReferDataID: %d,\n", indentationValues, dataStoreSearchParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreSearchParam.Tags))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %d,\n", indentationValues, dataStoreSearchParam.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %d,\n", indentationValues, dataStoreSearchParam.ResultOrder))

	if dataStoreSearchParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dataStoreSearchParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sResultOption: %d,\n", indentationValues, dataStoreSearchParam.ResultOption))
	b.WriteString(fmt.Sprintf("%sMinimalRatingFrequency: %v,\n", indentationValues, dataStoreSearchParam.MinimalRatingFrequency))
	b.WriteString(fmt.Sprintf("%sUseCache: %t,\n", indentationValues, dataStoreSearchParam.UseCache))
	b.WriteString(fmt.Sprintf("%sTotalCountEnabled: %t,\n", indentationValues, dataStoreSearchParam.TotalCountEnabled))
	b.WriteString(fmt.Sprintf("%sDataTypes: %v\n", indentationValues, dataStoreSearchParam.DataTypes))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	return &DataStoreSearchParam{}
}
