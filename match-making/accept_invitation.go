// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptInvitation sets the AcceptInvitation handler function
func (protocol *Protocol) AcceptInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32) {
	protocol.acceptInvitationHandler = handler
}

func (protocol *Protocol) handleAcceptInvitation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acceptInvitationHandler == nil {
		globals.Logger.Warning("MatchMaking::AcceptInvitation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.acceptInvitationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.acceptInvitationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.acceptInvitationHandler(nil, packet, callID, idGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
