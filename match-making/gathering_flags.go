package match_making

type gatheringFlags struct {
	DisconnectChangeOwner uint32 // TODO - Does this really only happen when a disconnect happens, or can the owner change at other times?
	Unknown1              uint32
	Unknown2              uint32
}

var GatheringFlags = gatheringFlags{
	DisconnectChangeOwner: 0x10,
	Unknown1:              0x20,
	Unknown2:              0x200,
}