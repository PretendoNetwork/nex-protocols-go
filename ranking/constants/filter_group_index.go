package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FilterGroupIndex is used by RankingOrderParam.GroupIndex to select which group to filter by in a score request
type FilterGroupIndex uint8

// WriteTo writes the FilterGroupIndex to the given writable
func (fgi FilterGroupIndex) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(fgi))
}

// ExtractFrom extracts the FilterGroupIndex value from the given readable
func (fgi *FilterGroupIndex) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*fgi = FilterGroupIndex(value)
	return nil
}

// String returns a human-readable representation of the FilterGroupIndex.
func (fgi FilterGroupIndex) String() string {
	switch fgi {
	case FilterGroupIndex0:
		return "Index0"
	case FilterGroupIndex1:
		return "Index1"
	case FilterGroupIndex2:
		return "Index2"
	case FilterGroupIndex3:
		return "Index3"
	case FilterGroupIndexNone:
		return "None"
	default:
		return fmt.Sprintf("FilterGroupIndex(%d)", int(fgi))
	}
}

const (
	// FilterGroupIndex0 indicates RankingOrderParam.GroupNum should be compared to the 0th group.
	FilterGroupIndex0 FilterGroupIndex = iota

	// FilterGroupIndex1 indicates RankingOrderParam.GroupNum should be compared to the 1st group.
	FilterGroupIndex1

	// FilterGroupIndex2 indicates RankingOrderParam.GroupNum should be compared to the 2nd group.
	FilterGroupIndex2

	// FilterGroupIndex3 indicates RankingOrderParam.GroupNum should be compared to the 3rd group.
	FilterGroupIndex3

	// FilterGroupIndexNone indicates that no group filtering should be performed.
	FilterGroupIndexNone = 255
)
