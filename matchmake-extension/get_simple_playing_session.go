// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetSimplePlayingSession == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimplePlayingSession not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	listPID := types.NewList[*types.PID]()
	listPID.Type = types.NewPID(0)
	err = listPID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	includeLoginUser := types.NewPrimitiveBool(false)
	err = includeLoginUser.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetSimplePlayingSession(nil, packet, callID, listPID, includeLoginUser)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
