// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetStats == nil {
		globals.Logger.Warning("MatchMaking::GetStats not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstParticipants, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read lstParticipants from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstColumns, err := parametersStream.ReadBuffer() // * This is documented as List<byte>, but that's justs a buffer so...
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read lstColumns from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetStats(nil, packet, callID, idGathering, lstParticipants, lstColumns)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
