package constants

// AnybodyParticipationPolicyArgument seems to determine whether or not to
// close participation when the gathering owner changes when MatchmakeSystemTypeAnybody
// is used?
type AnybodyParticipationPolicyArgument = PolicyArgument

// IsValid ensures the value of the AnybodyParticipationPolicyArgument is within
// the expected range
func (appa AnybodyParticipationPolicyArgument) IsValid() bool {
	return appa >= AnybodyParticipationPolicyArgumentWithoutClose && appa <= AnybodyParticipationPolicyArgumentCloseOnOwnerMigration
}

const (
	// AnybodyParticipationPolicyArgumentWithoutClose means that participation should not
	// close when a new owner is selected.
	AnybodyParticipationPolicyArgumentWithoutClose AnybodyParticipationPolicyArgument = iota

	// AnybodyParticipationPolicyArgumentCloseOnOwnerMigration means that participation should
	// close when a new owner is selected.
	AnybodyParticipationPolicyArgumentCloseOnOwnerMigration
)
