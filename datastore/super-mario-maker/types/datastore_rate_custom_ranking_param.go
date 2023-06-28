package datastore_super_mario_maker_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRateCustomRankingParam is sent in the RateCustomRanking method
type DataStoreRateCustomRankingParam struct {
	nex.Structure
	DataID        uint64
	ApplicationId uint32
	Score         uint32
	Period        uint16
}

// ExtractFromStream extracts a DataStoreRateCustomRankingParam structure from a stream
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRateCustomRankingParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.ApplicationId, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationId. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateCustomRankingParam()

	copied.DataID = dataStoreRateCustomRankingParam.DataID
	copied.ApplicationId = dataStoreRateCustomRankingParam.ApplicationId
	copied.Score = dataStoreRateCustomRankingParam.Score
	copied.Period = dataStoreRateCustomRankingParam.Period

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateCustomRankingParam)

	if dataStoreRateCustomRankingParam.DataID != other.DataID {
		return false
	}

	if dataStoreRateCustomRankingParam.ApplicationId != other.ApplicationId {
		return false
	}

	if dataStoreRateCustomRankingParam.Score != other.Score {
		return false
	}

	if dataStoreRateCustomRankingParam.Period != other.Period {
		return false
	}

	return true
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	return &DataStoreRateCustomRankingParam{}
}
