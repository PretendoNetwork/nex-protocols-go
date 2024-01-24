// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	matchmakingVersion := protocol.server.MatchMakingProtocolVersion()

	var errorCode uint32

	if protocol.CreateMatchmakeSession == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSession not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	anyGathering := types.NewAnyDataHolder()
	err := anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message := types.NewString("")
	err = message.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount := types.NewPrimitiveU16(0)

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		participationCountU16, err := parametersStream.ReadPrimitiveUInt16LE()
		if err != nil {
			_, errorCode = protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from participationCount. %s", err.Error()), packet, callID, nil, nil, nil)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}

		participationCount = types.NewPrimitiveU16(participationCountU16)
	}

	rmcMessage, errorCode := protocol.CreateMatchmakeSession(nil, packet, callID, anyGathering, message, participationCount)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
