// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByParticipants(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.FindByParticipants == nil {
		globals.Logger.Warning("MatchMaking::FindByParticipants not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewList[*types.PID]()
	pid.Type = types.NewPID(0)
	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByParticipants(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByParticipants(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
