// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	matchmakingVersion := protocol.server.MatchMakingProtocolVersion()

	var errorCode uint32

	if protocol.CreateMatchmakeSession == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSession not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	var participationCount uint16 = 0

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		participationCount, err = parametersStream.ReadUInt16LE()
		if err != nil {
			_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from participationCount. %s", err.Error()), packet, callID, nil, "", 0)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}
	}

	rmcMessage, errorCode := protocol.CreateMatchmakeSession(nil, packet, callID, anyGathering, message, participationCount)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
