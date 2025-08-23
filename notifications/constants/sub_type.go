package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// SubType exists solely to restrict the kinds of values that can be passed
// to SubType.Build()
type SubType uint32

// WriteTo writes the SubType to the given writable
func (st SubType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(st))
}

// ExtractFrom extracts the SubType value from the given readable
func (st *SubType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*st = SubType(value)
	return nil
}
