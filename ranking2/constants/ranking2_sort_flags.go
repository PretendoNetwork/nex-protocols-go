package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2SortFlags determines how ranking results should be ordered.
type Ranking2SortFlags uint32

// WriteTo writes the Ranking2SortFlags to the given writable
func (r2sf Ranking2SortFlags) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(r2sf))
}

// ExtractFrom extracts the Ranking2SortFlags value from the given readable
func (r2sf *Ranking2SortFlags) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*r2sf = Ranking2SortFlags(value)
	return nil
}

// HasFlag checks if a given flag is set
func (r2sf Ranking2SortFlags) HasFlag(flag Ranking2SortFlags) bool {
	return r2sf&flag == flag
}

// HasFlag checks if all given flags are set
func (r2sf Ranking2SortFlags) HasFlags(flags ...Ranking2SortFlags) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if r2sf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the Ranking2SortFlags bitmask.
// Multiple flags are joined with "|", e.g. "Ascending|MoveToTopInTie".
// Returns "Nothing" if no flags are set.
func (r2sf Ranking2SortFlags) String() string {
	if r2sf == Ranking2SortFlagsNothing {
		return "Nothing"
	}

	flags := []struct {
		flag Ranking2SortFlags
		name string
	}{
		{Ranking2SortFlagsAscending, "Ascending"},
		{Ranking2SortFlagsDescending, "Descending"},
		{Ranking2SortFlagsMoveToTopInTie, "MoveToTopInTie"},
	}

	var parts []string
	for _, f := range flags {
		if r2sf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// Ranking2SortFlagsNothing means results should be returned in no
	// specific order.
	Ranking2SortFlagsNothing Ranking2SortFlags = 0

	// Ranking2SortFlagsAscending means results should be returned in no
	// ascending order?
	//
	// Note: This is a guess based on some light test behavior and the fact
	// that there is no other ordering flag like this for the request it's
	// used in. The real name of this is unknown. And functionality not 100%
	// confirmed.
	Ranking2SortFlagsAscending Ranking2SortFlags = 1

	// Ranking2SortFlagsDescending means results should be returned in no
	// descending order?
	//
	// Note: This is a guess based on some light test behavior and the fact
	// that there is no other ordering flag like this for the request it's
	// used in. The real name of this is unknown. And functionality not 100%
	// confirmed.
	Ranking2SortFlagsDescending Ranking2SortFlags = 2

	// Ranking2SortFlagsMoveToTopInTie means in the event of a tie between
	// another user and the caller, the caller is ranked better than the
	// other user.
	Ranking2SortFlagsMoveToTopInTie Ranking2SortFlags = 4
)
