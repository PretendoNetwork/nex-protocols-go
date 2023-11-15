// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetSimplePlayingSession == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimplePlayingSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listPID, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	includeLoginUser, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), packet, callID, nil, false)
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
