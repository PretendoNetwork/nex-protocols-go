// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetStats == nil {
		globals.Logger.Warning("MatchMaking::GetStats not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstParticipants := types.NewList[*types.PID]()
	lstParticipants.Type = types.NewPID(0)
	err = lstParticipants.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read lstParticipants from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstColumns := types.NewBuffer(nil)
	err = lstColumns.ExtractFrom(parametersStream) // * This is documented as List<byte>, but that's justs a buffer so...
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read lstColumns from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
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
