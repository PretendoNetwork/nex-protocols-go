package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (protocol *MatchmakeExtensionProtocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)) {
	protocol.GetSimplePlayingSessionHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleGetSimplePlayingSession(packet nex.PacketInterface) {
	if protocol.GetSimplePlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimplePlayingSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listPID := parametersStream.ReadListUInt32LE()
	includeLoginUser := parametersStream.ReadUInt8() == 1

	go protocol.GetSimplePlayingSessionHandler(nil, client, callID, listPID, includeLoginUser)
}
