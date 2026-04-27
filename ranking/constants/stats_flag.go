package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// StatsFlag is a bitmask used by GetRanking to request the inclusion of different aggregate stats
type StatsFlag uint32

// WriteTo writes the StatsFlag to the given writable
func (sf StatsFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(sf))
}

// ExtractFrom extracts the StatsFlag value from the given readable
func (sf *StatsFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*sf = StatsFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (sf StatsFlag) HasFlag(flag StatsFlag) bool {
	return sf&flag == flag
}

// HasFlag checks if all given flags are set
func (sf StatsFlag) HasFlags(flags ...StatsFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if sf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the StatsFlag bitmask.
// Multiple flags are joined with "|", e.g. "Total|Sum|Average".
// Returns "None" if no flags are set.
func (sf StatsFlag) String() string {
	if sf == 0 {
		return "None"
	}

	flags := []struct {
		flag StatsFlag
		name string
	}{
		{StatsFlagTotal, "Total"},
		{StatsFlagSum, "Sum"},
		{StatsFlagMin, "Min"},
		{StatsFlagMax, "Max"},
		{StatsFlagAverage, "Average"},
	}

	var parts []string
	for _, f := range flags {
		if sf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// StatsFlagTotal requests the total of the stats
	StatsFlagTotal = 0x1

	// StatsFlagSum requests the sum of the stats
	StatsFlagSum = 0x2

	// StatsFlagMin requests the minimum stat
	StatsFlagMin = 0x4

	// StatsFlagMax requests the maximum stat
	StatsFlagMax = 0x8

	// StatsFlagAverage requests the average of the stats
	StatsFlagAverage = 0x10
)
