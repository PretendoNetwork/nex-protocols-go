package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RatingFlag indicates how the server should handle
// user ratings for object rating slots
type RatingFlag uint8

// WriteTo writes the RatingFlag to the given writable
func (rf RatingFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(rf))
}

// ExtractFrom extracts the RatingFlag value from the given readable
func (rf *RatingFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*rf = RatingFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (rf RatingFlag) HasFlag(flag RatingFlag) bool {
	return rf&flag == flag
}

// HasFlag checks if all given flags are set
func (rf RatingFlag) HasFlags(flags ...RatingFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if rf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the RatingFlag bitmask.
// Multiple flags are joined with "|", e.g. "Modifiable|DisableSelfRating".
// Returns "None" if no flags are set.
func (rf RatingFlag) String() string {
	if rf == 0 {
		return "None"
	}

	flags := []struct {
		flag RatingFlag
		name string
	}{
		{RatingFlagModifiable, "Modifiable"},
		{RatingFlagRoundMinus, "RoundMinus"},
		{RatingFlagDisableSelfRating, "DisableSelfRating"},
	}

	var parts []string
	for _, f := range flags {
		if rf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// RatingFlagModifiable means that if a user rates an object
	// rating slot more than once, update the existing rating
	// rather than create a new one. If this flag is set, ratings
	// to slots with existing ratings from the user do not count
	// towards the rating count. If this flag is not set, ratings
	// to slots with existing ratings from the user are treated
	// as separate ratings and the rating count is incremented
	RatingFlagModifiable RatingFlag = 0x4

	// RatingFlagRoundMinus means that any rating value smaller
	// than 0 is rounded up to 0 before being handled
	RatingFlagRoundMinus RatingFlag = 0x8

	// RatingFlagDisableSelfRating means that a user cannot rate
	// an object that they own. If this flag is set and a user
	// tries to rate an object they own, DataStore::OperationNotAllowed is thrown
	RatingFlagDisableSelfRating RatingFlag = 0x10
)
