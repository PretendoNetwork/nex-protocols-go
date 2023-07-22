// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// OpenParticipation sets the OpenParticipation handler function
func (protocol *MatchmakeExtensionProtocol) OpenParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.openParticipationHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleOpenParticipation(packet nex.PacketInterface) {
	if protocol.openParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::OpenParticipation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.openParticipationHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.openParticipationHandler(nil, client, callID, gid)
}
