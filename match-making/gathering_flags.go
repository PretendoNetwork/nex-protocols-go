// Package protocol implements the Match Making protocol
package protocol

type gatheringFlags struct {
	PersistentGathering                   uint32
	DisconnectChangeOwner                 uint32
	Unknown1                              uint32
	PersistentGatheringLeaveParticipation uint32
	PersistentGatheringAllowZeroUsers     uint32
	ParticipantsChangeOwner               uint32
	VerboseParticipants                   uint32
	VerboseParticipantsEx                 uint32
}

// GatheringFlags is an enum of the possible flags for a gathering
var GatheringFlags = gatheringFlags{
	PersistentGathering:                   0x1,
	DisconnectChangeOwner:                 0x10,
	Unknown1:                              0x20,
	PersistentGatheringLeaveParticipation: 0x40,
	PersistentGatheringAllowZeroUsers:     0x80,
	ParticipantsChangeOwner:               0x200,
	VerboseParticipants:                   0x400,
	VerboseParticipantsEx:                 0x800,
}
