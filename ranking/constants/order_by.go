package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// OrderBy is used in RankingScoreData.OrderBy to set the "golf scoring" mode for a category.
type OrderBy uint8

const (
	OrderByAscending OrderBy = iota
	OrderByDescending
)

// WriteTo writes the OrderBy to the given writable
func (ob OrderBy) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(ob))
}

// ExtractFrom extracts the OrderBy value from the given readable
func (ob *OrderBy) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*ob = OrderBy(value)
	return nil
}
