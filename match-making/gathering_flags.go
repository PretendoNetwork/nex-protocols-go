// Package protocol implements the Match Making protocol
package protocol

type gatheringFlags struct {
	DisconnectChangeOwner *types.PrimitiveU32 // TODO - Does this really only happen when a disconnect happens, or can the owner change at other times?
	Unknown1              *types.PrimitiveU32
	Unknown2              *types.PrimitiveU32
}

// GatheringFlags is an enum of the possible flags for a gathering
var GatheringFlags = gatheringFlags{
	DisconnectChangeOwner: 0x10,
	Unknown1:              0x20,
	Unknown2:              0x200,
}
