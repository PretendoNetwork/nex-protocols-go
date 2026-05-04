package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RatingInternalFlag indicates whether or not a minimum or
// maximum rating value should be set
type RatingInternalFlag uint8

// WriteTo writes the RatingInternalFlag to the given writable
func (rif RatingInternalFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(rif))
}

// ExtractFrom extracts the RatingInternalFlag value from the given readable
func (rif *RatingInternalFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*rif = RatingInternalFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (rif RatingInternalFlag) HasFlag(flag RatingInternalFlag) bool {
	return rif&flag == flag
}

// HasFlag checks if all given flags are set
func (rif RatingInternalFlag) HasFlags(flags ...RatingInternalFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if rif&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the RatingInternalFlag bitmask.
// Multiple flags are joined with "|", e.g. "UseRangeMin|UseRangeMax".
// Returns "None" if no flags are set.
func (rif RatingInternalFlag) String() string {
	if rif == 0 {
		return "None"
	}

	flags := []struct {
		flag RatingInternalFlag
		name string
	}{
		{RatingInternalFlagUseRangeMin, "UseRangeMin"},
		{RatingInternalFlagUseRangeMax, "UseRangeMax"},
	}

	var parts []string
	for _, f := range flags {
		if rif&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// RatingInternalFlagUseRangeMin means that the DataStoreRatingInitParam.rangeMin
	// value should be respected. If this flag is not set, the minimum value is ignored.
	// If this flag is set and a rating is below the minimum, `DataStore::InvalidArgument`
	// is thrown
	RatingInternalFlagUseRangeMin RatingInternalFlag = 0x2

	// RatingInternalFlagUseRangeMax means that the DataStoreRatingInitParam.rangeMax
	// value should be respected. If this flag is not set, the maximum value is ignored.
	// If this flag is set and a rating is above the maximum, `DataStore::InvalidArgument`
	// is thrown
	RatingInternalFlagUseRangeMax RatingInternalFlag = 0x4
)
