package constants

import (
	"fmt"

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

// String returns a human-readable representation of the ParticipationPolicy.
func (pp ParticipationPolicy) String() string {
	switch pp {
	case ParticipationPolicyPasswordProtected:
		return "PasswordProtected"
	case ParticipationPolicyOpenParticipation:
		return "OpenParticipation"
	case ParticipationPolicyAnybody:
		return "Anybody"
	case ParticipationPolicyInviteOnly:
		return "InviteOnly"
	case ParticipationPolicyCommunity:
		return "Community"
	case ParticipationPolicyFriendsOnly:
		return "FriendsOnly"
	default:
		return fmt.Sprintf("ParticipationPolicy(%d)", int(pp))
	}
}

const (
	// ParticipationPolicyPasswordProtected indicates that a session is protected by as password.
	//
	// When used, the PolicyArgument is the first four bytes of the MD5 hash of the key given
	// to the participant
	ParticipationPolicyPasswordProtected ParticipationPolicy = 4

	// ParticipationPolicyOpenParticipation indicates that a session is open to anyone.
	ParticipationPolicyOpenParticipation ParticipationPolicy = 8

	// ParticipationPolicyAnybody seems to function the exact same
	// as ParticipationPolicyOpenParticipation.
	ParticipationPolicyAnybody ParticipationPolicy = 95

	// ParticipationPolicyInviteOnly has an unknown function. This is used when MatchmakeSystemTypeInvite
	// is used, and thus is assumed to be invitation-based, but not limited to friend invitations.
	ParticipationPolicyInviteOnly ParticipationPolicy = 96

	// ParticipationPolicyCommunity indicates that a session is a community/persistent
	// gathering.
	//
	// Seen in Mario Kart 7, Mario Golf, etc.
	ParticipationPolicyCommunity ParticipationPolicy = 97

	// ParticipationPolicyFriendsOnly indicates that only friends may participate.
	ParticipationPolicyFriendsOnly ParticipationPolicy = 98
)
