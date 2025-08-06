package constants

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ParticipationPolicy indicates the session participation policy.
//
// Note: We do not know the real names for any of these, this is
// all guess work.
type ParticipationPolicy uint32

// WriteTo writes the ParticipationPolicy to the given writable
func (pp ParticipationPolicy) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(pp))
}

// ExtractFrom extracts the ParticipationPolicy value from the given readable
func (pp *ParticipationPolicy) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*pp = ParticipationPolicy(value)
	return nil
}

const (
	// ParticipationPolicyOpenParticipation indicates that a session is open to anyone.
	ParticipationPolicyOpenParticipation ParticipationPolicy = 8

	// ParticipationPolicyNintendoOpenParticipation seems to function the exact same
	// as ParticipationPolicyOpenParticipation.
	ParticipationPolicyNintendoOpenParticipation ParticipationPolicy = 95

	// ParticipationPolicyCommunity indicates that a session is a community//persistent
	// gathering.
	//
	// Seen in Mario Kart 7, Mario Golf, etc.
	ParticipationPolicyCommunity ParticipationPolicy = 97

	// ParticipationPolicyFriendsOnly indicates that only friends may participate.
	ParticipationPolicyFriendsOnly ParticipationPolicy = 98
)
