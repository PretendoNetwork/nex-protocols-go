// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetDetailedParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetDetailedParticipants == nil {
		globals.Logger.Warning("MatchMakingExt::GetDetailedParticipants not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetDetailedParticipants(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bOnlyActive, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.GetDetailedParticipants(fmt.Errorf("Failed to read bOnlyActive from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetDetailedParticipants(nil, packet, callID, idGathering, bOnlyActive)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
