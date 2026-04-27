package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeGeoIPResult represents an enum with an unknown use
type MatchmakeGeoIPResult uint32

// WriteTo writes the MatchmakeGeoIPResult to the given writable
func (mgir MatchmakeGeoIPResult) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(mgir))
}

// ExtractFrom extracts the MatchmakeGeoIPResult value from the given readable
func (mgir *MatchmakeGeoIPResult) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*mgir = MatchmakeGeoIPResult(value)
	return nil
}

// String returns a human-readable representation of the MatchmakeGeoIPResult.
func (mgir MatchmakeGeoIPResult) String() string {
	switch mgir {
	case MatchmakeGeoIPResultInvalid:
		return "Invalid"
	case MatchmakeGeoIPResultFound:
		return "Found"
	case MatchmakeGeoIPResultNotFound:
		return "NotFound"
	case MatchmakeGeoIPResultUnused:
		return "Unused"
	default:
		return fmt.Sprintf("MatchmakeGeoIPResult(%d)", int(mgir))
	}
}

const (
	// MatchmakeGeoIPResultInvalid indicates an invalid value
	MatchmakeGeoIPResultInvalid MatchmakeGeoIPResult = iota

	// MatchmakeGeoIPResultAnybody has an unknown use
	MatchmakeGeoIPResultFound

	// MatchmakeGeoIPResultFriends has an unknown use
	MatchmakeGeoIPResultNotFound

	// MatchmakeGeoIPResultFriendsInvite has an unknown use
	MatchmakeGeoIPResultUnused
)
