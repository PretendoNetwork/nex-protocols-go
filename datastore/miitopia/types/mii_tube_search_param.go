// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

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

// ExtractFrom extracts the DataStoreGetMetaByOwnerIDParam from the given readable
func (mtsp *MiiTubeSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mtsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam header. %s", err.Error())
	}

	err = mtsp.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.Name. %s", err.Error())
	}

	err = mtsp.Page.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.Page. %s", err.Error())
	}

	err = mtsp.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.Category. %s", err.Error())
	}

	err = mtsp.Gender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.Gender. %s", err.Error())
	}

	err = mtsp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.Country. %s", err.Error())
	}

	err = mtsp.SearchType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.SearchType. %s", err.Error())
	}

	err = mtsp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption. %s", err.Error())
	}

	return nil
}

// NewMiiTubeSearchParam returns a new MiiTubeSearchParam
func NewMiiTubeSearchParam() *MiiTubeSearchParam {
	mtsp := &MiiTubeSearchParam{
		Name:         types.NewString(""),
		Page:         types.NewPrimitiveU32(0),
		Category:     types.NewPrimitiveU8(0),
		Gender:       types.NewPrimitiveU8(0),
		Country:      types.NewPrimitiveU8(0),
		SearchType:   types.NewPrimitiveU8(0),
		ResultOption: types.NewPrimitiveU8(0),
	}

	return mtsp
}
