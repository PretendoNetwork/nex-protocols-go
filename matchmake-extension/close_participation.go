package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CloseParticipation sets the CloseParticipation handler function
func (protocol *MatchmakeExtensionProtocol) CloseParticipation(handler func(err error, client *nex.Client, callID uint32, GID uint32)) {
	protocol.CloseParticipationHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleCloseParticipation(packet nex.PacketInterface) {
	if protocol.CloseParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CloseParticipation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	GID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.CloseParticipationHandler(fmt.Errorf("Failed to read GID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.CloseParticipationHandler(nil, client, callID, GID)
}
