package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// OrderBy is used in RankingScoreData.OrderBy to set the "golf scoring" mode for a category.
type OrderBy uint8

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

// String returns a human-readable representation of the OrderBy.
func (ob OrderBy) String() string {
	switch ob {
	case OrderByAscending:
		return "Ascending"
	case OrderByDescending:
		return "Descending"
	default:
		return fmt.Sprintf("OrderBy(%d)", int(ob))
	}
}

const (
	// OrderByAscending indicates sorting scores in ascending order
	OrderByAscending OrderBy = iota

	// OrderByDescending indicates sorting scores in descending order
	OrderByDescending
)
