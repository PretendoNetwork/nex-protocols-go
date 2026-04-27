package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlatformType denotes the type of console platform the client
// is using.
//
// Note: This is not a real type. This is a bespoke type created
// for our convenience.
type PlatformType uint32

// WriteTo writes the PlatformType to the given writable
func (pt PlatformType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(pt))
}

// ExtractFrom extracts the PlatformType value from the given readable
func (pt *PlatformType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*pt = PlatformType(value)
	return nil
}

// String returns a human-readable representation of the PlatformType.
func (pt PlatformType) String() string {
	switch pt {
	case PlatformType3DS:
		return "3DS"
	case PlatformTypeWiiU:
		return "WiiU"
	case PlatformTypeSwitch:
		return "Switch"
	default:
		return fmt.Sprintf("PlatformType(%d)", int(pt))
	}
}

const (
	// PlatformType3DS means that the connecting client is on a 3DS.
	PlatformType3DS PlatformType = iota + 1

	// PlatformTypeWiiU means that the connecting client is on a Wii U.
	PlatformTypeWiiU

	// PlatformTypeSwitch means that the connecting client is on a Nintendo Switch.
	PlatformTypeSwitch
)
