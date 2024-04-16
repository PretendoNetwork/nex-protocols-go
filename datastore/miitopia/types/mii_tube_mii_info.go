// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// MiiTubeMiiInfo is a type within the DataStoreMiitopia protocol
type MiiTubeMiiInfo struct {
	types.Structure
	MetaInfo    *datastore_types.DataStoreMetaInfo
	Category    *types.PrimitiveU8
	RankingType *types.PrimitiveU8
}

func NewMiiTubeMiiInfo() *MiiTubeMiiInfo {
	mtmi := &MiiTubeMiiInfo{
		MetaInfo:    datastore_types.NewDataStoreMetaInfo(),
		Category:    types.NewPrimitiveU8(0),
		RankingType: types.NewPrimitiveU8(0),
	}

	return mtmi
}
