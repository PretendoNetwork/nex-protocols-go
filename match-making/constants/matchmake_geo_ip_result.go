package constants

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeGeoIPResult represents an enum with an unknown use
type MatchmakeGeoIPResult uint32

// WriteTo writes the MatchmakeGeoIPResult to the given writable
func (mgipr MatchmakeGeoIPResult) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(mgipr))
}

// ExtractFrom extracts the MatchmakeGeoIPResult value from the given readable
func (mgipr *MatchmakeGeoIPResult) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*mgipr = MatchmakeGeoIPResult(value)
	return nil
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
