package constants

// AnybodyParticipationPolicyArgument represents the participation policy
// arguments for when ParticipationPolicyAnybody is used.
type AnybodyParticipationPolicyArgument = PolicyArgument

const (
	// AnybodyParticipationPolicyArgumentWithoutClose means that participation should not
	// close when a new owner is selected.
	AnybodyParticipationPolicyArgumentWithoutClose AnybodyParticipationPolicyArgument = iota

	// AnybodyParticipationPolicyArgumentCloseOnOwnerMigration means that participation should
	// close when a new owner is selected.
	AnybodyParticipationPolicyArgumentCloseOnOwnerMigration
)
