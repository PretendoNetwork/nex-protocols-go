package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeSystemType represents the method of matchmaking being used
type MatchmakeSystemType uint32

// WriteTo writes the MatchmakeSystemType to the given writable
func (mst MatchmakeSystemType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(mst))
}

// ExtractFrom extracts the MatchmakeSystemType value from the given readable
func (mst *MatchmakeSystemType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*mst = MatchmakeSystemType(value)
	if !mst.IsValid() {
		return fmt.Errorf("Value %d is out of range", *mst)
	}

	return nil
}

// IsValid ensures the value of the MatchmakeSystemType is within
// the expected range
func (mst MatchmakeSystemType) IsValid() bool {
	return mst >= MatchmakeSystemTypeInvalid && mst <= MatchmakeSystemTypePersistentGathering
}

const (
	// MatchmakeSystemTypeInvalid indicates an invalid value
	MatchmakeSystemTypeInvalid MatchmakeSystemType = iota

	// MatchmakeSystemTypeAnybody indicates that anybody can join the MatchmakeSession
	MatchmakeSystemTypeAnybody

	// MatchmakeSystemTypeFriends indicates that only friends of the owner can join the MatchmakeSession
	MatchmakeSystemTypeFriends

	// MatchmakeSystemTypeFriendsInvite indicates that only friends of the owner with invitation can join the MatchmakeSession
	MatchmakeSystemTypeFriendsInvite

	// MatchmakeSystemTypeFriends indicates that anybody with an invitation can join the MatchmakeSession
	MatchmakeSystemTypeInvite

	// MatchmakeSystemTypeFriends indicates that the MatchmakeSession is part of a PersistentGathering
	MatchmakeSystemTypePersistentGathering
)
