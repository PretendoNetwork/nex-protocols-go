// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Invite sets the Invite handler function
func (protocol *Protocol) Invite(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string) uint32) {
	protocol.inviteHandler = handler
}

func (protocol *Protocol) handleInvite(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.inviteHandler == nil {
		globals.Logger.Warning("MatchMaking::Invite not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.inviteHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstPrincipals, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.inviteHandler(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.inviteHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.inviteHandler(nil, packet, callID, idGathering, lstPrincipals, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
