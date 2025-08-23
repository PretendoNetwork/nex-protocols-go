package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// Ranking2GetOptionFlags determines what data is returned when requesting common data.
type Ranking2GetOptionFlags uint32

const (
	// Ranking2GetOptionFlagsNothing means that no extra data should be returned
	// in the common data.
	//
	// Note: The `userName` and `binaryData` fields seem to always be populated, regardless
	// of flags. These seem to exist solely to enable the `mii` field.
	Ranking2GetOptionFlagsNothing Ranking2GetOptionFlags = 0

	// Ranking2GetOptionFlagsMii means that Mii data should be returned
	// in the common data.
	//
	// Note: This is a guess based on some light test behavior. The real
	// name of this is unknown.
	Ranking2GetOptionFlagsMii Ranking2GetOptionFlags = 1
)

// WriteTo writes the Ranking2GetOptionFlags to the given writable
func (r2gof Ranking2GetOptionFlags) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(r2gof))
}

// ExtractFrom extracts the Ranking2GetOptionFlags value from the given readable
func (r2gof *Ranking2GetOptionFlags) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*r2gof = Ranking2GetOptionFlags(value)
	return nil
}
