// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import "github.com/PretendoNetwork/nex-go/v2/types"

// MiiTubeSearchParam is a type within the DataStoreMiitopia protocol
type MiiTubeSearchParam struct {
	types.Structure
	Name         *types.String
	Page         *types.PrimitiveU32
	Category     *types.PrimitiveU8
	Gender       *types.PrimitiveU8
	Country      *types.PrimitiveU8
	SearchType   *types.PrimitiveU8
	ResultOption *types.PrimitiveU8
}
