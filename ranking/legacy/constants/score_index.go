package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// ScoreIndex is used to select the index for the score list to use when sorting.
//
// Note: The names of this type and its values are guesses based on context
type ScoreIndex uint8

// WriteTo writes the ScoreIndex to the given writable
func (si ScoreIndex) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(si))
}

// ExtractFrom extracts the ScoreIndex value from the given readable
func (si *ScoreIndex) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*si = ScoreIndex(value)
	return nil
}

const (
	// ScoreIndex0 indicates sorting by the 0th element on the score list
	ScoreIndex0 ScoreIndex = iota

	// ScoreIndex1 indicates sorting by the 1st element on the score list
	ScoreIndex1
)
