package constants

// MatchmakeSessionModificationFlag indicates the flags set on a gathering
type MatchmakeSessionModificationFlag uint32

const (
	// MatchmakeSessionModificationFlagNone modifies nothing.
	MatchmakeSessionModificationFlagNone MatchmakeSessionModificationFlag = 0x0

	// MatchmakeSessionModificationFlagAttributes updates the sessions attributes.
	MatchmakeSessionModificationFlagAttributes MatchmakeSessionModificationFlag = 0x1

	// MatchmakeSessionModificationFlagOpenParticipation updates the sessions open participation status.
	MatchmakeSessionModificationFlagOpenParticipation MatchmakeSessionModificationFlag = 0x2

	// MatchmakeSessionModificationFlagApplicationBuffer updates the sessions application buffer.
	MatchmakeSessionModificationFlagApplicationBuffer MatchmakeSessionModificationFlag = 0x4

	// MatchmakeSessionModificationFlagProgressScore updates the sessions progress score.
	MatchmakeSessionModificationFlagProgressScore MatchmakeSessionModificationFlag = 0x8

	// MatchmakeSessionModificationFlagOption0 updates the sessions option0 value.
	MatchmakeSessionModificationFlagOption0 MatchmakeSessionModificationFlag = 0x10

	// MatchmakeSessionModificationFlagMatchmakeParam updates the sessions MatchmakeParam.
	//
	// This flag will entirely replace the old MatchmakeParam.
	//
	// Cannot be used with MatchmakeSessionModificationFlagMatchmakeParamOverride.
	MatchmakeSessionModificationFlagMatchmakeParam MatchmakeSessionModificationFlag = 0x20

	// MatchmakeSessionModificationFlagMatchmakeParamOverride updates the sessions
	//
	// This flag will only update existing parameters and add new ones, but does not delete any.
	//
	// Cannot be used with MatchmakeSessionModificationFlagMatchmakeParam.
	MatchmakeSessionModificationFlagMatchmakeParamOverride MatchmakeSessionModificationFlag = 0x40

	// MatchmakeSessionModificationFlagStartedTime updates the sessions started time.
	MatchmakeSessionModificationFlagStartedTime MatchmakeSessionModificationFlag = 0x80

	// MatchmakeSessionModificationFlagUserPassword updates the sessions user password.
	MatchmakeSessionModificationFlagUserPassword MatchmakeSessionModificationFlag = 0x100

	// MatchmakeSessionModificationFlagGameMode updates the sessions game mode.
	MatchmakeSessionModificationFlagGameMode MatchmakeSessionModificationFlag = 0x200

	// MatchmakeSessionModificationFlagDescription updates the sessions description.
	MatchmakeSessionModificationFlagDescription MatchmakeSessionModificationFlag = 0x400

	// MatchmakeSessionModificationFlagMinParticipants updates the sessions minimum number of participants.
	MatchmakeSessionModificationFlagMinParticipants MatchmakeSessionModificationFlag = 0x800

	// MatchmakeSessionModificationFlagMaxParticipants updates the sessions maximum number of participants.
	MatchmakeSessionModificationFlagMaxParticipants MatchmakeSessionModificationFlag = 0x1000

	// MatchmakeSessionModificationFlagMatchmakeSystemType updates the sessions MatchmakeSystemType
	MatchmakeSessionModificationFlagMatchmakeSystemType MatchmakeSessionModificationFlag = 0x2000

	// MatchmakeSessionModificationFlagCodeword updates the sessions codeword
	MatchmakeSessionModificationFlagCodeword MatchmakeSessionModificationFlag = 0x4000
)
