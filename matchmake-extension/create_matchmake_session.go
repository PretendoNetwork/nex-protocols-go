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

	if protocol.CreateMatchmakeSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::CreateMatchmakeSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	anyGathering := types.NewAnyDataHolder()
	err := anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	message := types.NewString("")
	err = message.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	participationCount := types.NewPrimitiveU16(0)

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		participationCountU16, err := parametersStream.ReadPrimitiveUInt16LE()
		if err != nil {
			_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read message from participationCount. %s", err.Error()), packet, callID, nil, nil, nil)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}

		participationCount = types.NewPrimitiveU16(participationCountU16)
	}

	rmcMessage, rmcError := protocol.CreateMatchmakeSession(nil, packet, callID, anyGathering, message, participationCount)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
