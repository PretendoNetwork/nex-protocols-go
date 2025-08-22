package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// StatsFlag is a bitmask used by GetRanking to request the inclusion of different aggregate stats
type StatsFlag uint32

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
