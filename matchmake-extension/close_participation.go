// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CloseParticipation sets the CloseParticipation handler function
func (protocol *Protocol) CloseParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32) {
	protocol.closeParticipationHandler = handler
}

func (protocol *Protocol) handleCloseParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.closeParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CloseParticipation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.closeParticipationHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.closeParticipationHandler(nil, packet, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
