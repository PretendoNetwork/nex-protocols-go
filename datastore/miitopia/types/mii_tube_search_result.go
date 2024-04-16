// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiTubeSearchResult is a type within the DataStoreMiitopia protocol
type MiiTubeSearchResult struct {
	types.Structure

	Result  *types.List[*MiiTubeMiiInfo]
	Count   *types.PrimitiveU32
	Page    *types.PrimitiveU32
	HasNext *types.PrimitiveBool
}
