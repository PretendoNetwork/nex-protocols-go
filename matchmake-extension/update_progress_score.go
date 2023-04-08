package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProgressScore sets the UpdateProgressScore handler function
func (protocol *MatchmakeExtensionProtocol) UpdateProgressScore(handler func(err error, client *nex.Client, callID uint32, GID uint32, progressScore uint8)) {
	protocol.UpdateProgressScoreHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleUpdateProgressScore(packet nex.PacketInterface) {
	if protocol.UpdateProgressScoreHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateProgressScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	GID := parametersStream.ReadUInt32LE()
	progressScore := parametersStream.ReadUInt8()

	go protocol.UpdateProgressScoreHandler(nil, client, callID, GID, progressScore)
}
