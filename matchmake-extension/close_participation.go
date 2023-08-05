// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CloseParticipation sets the CloseParticipation handler function
func (protocol *Protocol) CloseParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32) uint32) {
	protocol.closeParticipationHandler = handler
}

func (protocol *Protocol) handleCloseParticipation(packet nex.PacketInterface) {
	if protocol.closeParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CloseParticipation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.closeParticipationHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.closeParticipationHandler(nil, client, callID, gid)
}
