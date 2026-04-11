package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// PolicyArgument is not an official type. It exists to be the
// base type of other policy argument types and to be used as
// the Gathering.PolicyArgument field type.
type PolicyArgument uint32

// WriteTo writes the PolicyArgument to the given writable
func (pa PolicyArgument) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(pa))
}

// ExtractFrom extracts the PolicyArgument value from the given readable
func (pa *PolicyArgument) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*pa = PolicyArgument(value)
	return nil
}
