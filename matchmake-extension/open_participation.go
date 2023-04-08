package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// OpenParticipation sets the OpenParticipation handler function
func (protocol *MatchmakeExtensionProtocol) OpenParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.OpenParticipationHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleOpenParticipation(packet nex.PacketInterface) {
	if protocol.OpenParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::OpenParticipation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	gid := parametersStream.ReadUInt32LE()

	go protocol.OpenParticipationHandler(nil, client, callID, gid)
}
