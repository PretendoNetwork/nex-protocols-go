// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPlayingSession(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetPlayingSession == nil {
		globals.Logger.Warning("MatchmakeExtension::GetPlayingSession not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstPID := types.NewList[*types.PID]()
	lstPID.Type = types.NewPID(0)
	err = lstPID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetPlayingSession(fmt.Errorf("Failed to read lstPID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPlayingSession(nil, packet, callID, lstPID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
