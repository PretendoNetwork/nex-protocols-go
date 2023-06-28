package datastore_super_mario_maker_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCustomRankingByDataIdParam is sent in the GetCustomRankingByDataId method
type DataStoreGetCustomRankingByDataIdParam struct {
	nex.Structure
	ApplicationId uint32
	DataIdList    []uint64
	ResultOption  uint8
}

// ExtractFromStream extracts a DataStoreGetCustomRankingByDataIdParam structure from a stream
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetCustomRankingByDataIdParam.ApplicationId, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.ApplicationId. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIdParam.DataIdList, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.DataIdList. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIdParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.ResultOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCustomRankingByDataIdParam
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCustomRankingByDataIdParam()

	copied.ApplicationId = dataStoreGetCustomRankingByDataIdParam.ApplicationId
	copied.DataIdList = make([]uint64, len(dataStoreGetCustomRankingByDataIdParam.DataIdList))

	copy(copied.DataIdList, dataStoreGetCustomRankingByDataIdParam.DataIdList)

	copied.ResultOption = dataStoreGetCustomRankingByDataIdParam.ResultOption

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCustomRankingByDataIdParam)

	if dataStoreGetCustomRankingByDataIdParam.ApplicationId != other.ApplicationId {
		return false
	}

	if len(dataStoreGetCustomRankingByDataIdParam.DataIdList) != len(other.DataIdList) {
		return false
	}

	for i := 0; i < len(dataStoreGetCustomRankingByDataIdParam.DataIdList); i++ {
		if dataStoreGetCustomRankingByDataIdParam.DataIdList[i] != other.DataIdList[i] {
			return false
		}
	}

	if dataStoreGetCustomRankingByDataIdParam.ResultOption != other.ResultOption {
		return false
	}

	return true
}

// NewDataStoreGetCustomRankingByDataIdParam returns a new DataStoreGetCustomRankingByDataIdParam
func NewDataStoreGetCustomRankingByDataIdParam() *DataStoreGetCustomRankingByDataIdParam {
	return &DataStoreGetCustomRankingByDataIdParam{}
}
