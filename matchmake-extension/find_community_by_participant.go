package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByParticipant sets the FindCommunityByParticipant handler function
func (protocol *MatchmakeExtensionProtocol) FindCommunityByParticipant(handler func(err error, client *nex.Client, callID uint32, pid uint32, resultRange *nex.ResultRange)) {
	protocol.FindCommunityByParticipantHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleFindCommunityByParticipant(packet nex.PacketInterface) {
	if protocol.FindCommunityByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByParticipant not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid := parametersStream.ReadUInt32LE()
	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.FindCommunityByParticipantHandler(nil, client, callID, 0, nil)
		return
	}

	go protocol.FindCommunityByParticipantHandler(nil, client, callID, pid, resultRange.(*nex.ResultRange))
}
