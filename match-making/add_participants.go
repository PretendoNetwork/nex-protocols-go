// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddParticipants == nil {
		globals.Logger.Warning("MatchMaking::AddParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.AddParticipants(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstPrincipals, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.AddParticipants(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.AddParticipants(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddParticipants(nil, packet, callID, idGathering, lstPrincipals, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
