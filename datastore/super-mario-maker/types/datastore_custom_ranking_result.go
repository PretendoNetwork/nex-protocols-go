package datastore_super_mario_maker_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCustomRankingResult is sent in the FollowingsLatestCourseSearchObject method
type DataStoreCustomRankingResult struct {
	nex.Structure
	Order    uint32
	Score    uint32
	MetaInfo *datastore_types.DataStoreMetaInfo
}

// ExtractFromStream extracts a DataStoreCustomRankingResult structure from a stream
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCustomRankingResult.Order, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Order. %s", err.Error())
	}

	dataStoreCustomRankingResult.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Score. %s", err.Error())
	}

	metaInfo, err := stream.ReadStructure(datastore_types.NewDataStoreMetaInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.MetaInfo. %s", err.Error())
	}

	dataStoreCustomRankingResult.MetaInfo = metaInfo.(*datastore_types.DataStoreMetaInfo)

	return nil
}

// Bytes encodes the DataStoreCustomRankingResult and returns a byte array
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Order)
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Score)
	stream.WriteStructure(dataStoreCustomRankingResult.MetaInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCustomRankingResult
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Copy() nex.StructureInterface {
	copied := NewDataStoreCustomRankingResult()

	copied.Order = dataStoreCustomRankingResult.Order
	copied.Score = dataStoreCustomRankingResult.Score
	copied.MetaInfo = dataStoreCustomRankingResult.MetaInfo.Copy().(*datastore_types.DataStoreMetaInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCustomRankingResult)

	if dataStoreCustomRankingResult.Order != other.Order {
		return false
	}

	if dataStoreCustomRankingResult.Score != other.Score {
		return false
	}

	if !dataStoreCustomRankingResult.MetaInfo.Equals(other.MetaInfo) {
		return false
	}

	return true
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{}
}
