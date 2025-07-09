package constants

// AnybodyParticipationPolicyArgument seems to determine whether or not to
// close participation when the gathering owner changes when MatchmakeSystemTypeAnybody
// is used.
type AnybodyParticipationPolicyArgument uint32

const (
	// AnybodyParticipationPolicyArgumentWithoutClose means that participation should not
	// close when a new host is selected.
	AnybodyParticipationPolicyArgumentWithoutClose AnybodyParticipationPolicyArgument = iota

	// AnybodyParticipationPolicyArgumentCloseOnOwnerMigration means that participation should
	// close when a new host is selected.
	AnybodyParticipationPolicyArgumentCloseOnOwnerMigration
)
