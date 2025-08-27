package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

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
