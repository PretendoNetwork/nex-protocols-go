package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	nex.Structure
	SearchTarget           uint8
	OwnerIds               []uint32
	OwnerType              uint8
	DestinationIds         []uint64
	DataType               uint16
	CreatedAfter           *nex.DateTime
	CreatedBefore          *nex.DateTime
	UpdatedAfter           *nex.DateTime
	UpdatedBefore          *nex.DateTime
	ReferDataId            uint32
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

	dataStoreSearchParam.OwnerIds, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerIds. %s", err.Error())
	}

	dataStoreSearchParam.OwnerType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerType. %s", err.Error())
	}

	dataStoreSearchParam.DestinationIds, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DestinationIds. %s", err.Error())
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

	dataStoreSearchParam.ReferDataId, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ReferDataId. %s", err.Error())
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
	copied.OwnerIds = make([]uint32, len(dataStoreSearchParam.OwnerIds))

	copy(copied.OwnerIds, dataStoreSearchParam.OwnerIds)

	copied.OwnerType = dataStoreSearchParam.OwnerType
	copied.DestinationIds = make([]uint64, len(dataStoreSearchParam.DestinationIds))

	copy(copied.DestinationIds, dataStoreSearchParam.DestinationIds)

	copied.DataType = dataStoreSearchParam.DataType
	copied.CreatedAfter = dataStoreSearchParam.CreatedAfter.Copy()
	copied.CreatedBefore = dataStoreSearchParam.CreatedBefore.Copy()
	copied.UpdatedAfter = dataStoreSearchParam.UpdatedAfter.Copy()
	copied.UpdatedBefore = dataStoreSearchParam.UpdatedBefore.Copy()
	copied.ReferDataId = dataStoreSearchParam.ReferDataId
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

	if len(dataStoreSearchParam.OwnerIds) != len(other.OwnerIds) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.OwnerIds); i++ {
		if dataStoreSearchParam.OwnerIds[i] != other.OwnerIds[i] {
			return false
		}
	}

	if dataStoreSearchParam.OwnerType != other.OwnerType {
		return false
	}

	if len(dataStoreSearchParam.DestinationIds) != len(other.DestinationIds) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.DestinationIds); i++ {
		if dataStoreSearchParam.DestinationIds[i] != other.DestinationIds[i] {
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

	if dataStoreSearchParam.ReferDataId != other.ReferDataId {
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

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	return &DataStoreSearchParam{}
}
