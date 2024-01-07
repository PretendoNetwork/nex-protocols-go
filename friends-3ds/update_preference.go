// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePreference == nil {
		globals.Logger.Warning("Friends3DS::UpdatePreference not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	publicMode := types.NewPrimitiveBool(false)
	err = publicMode.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read publicMode from parameters. %s", err.Error()), packet, callID, false, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showGame := types.NewPrimitiveBool(false)
	err = showGame.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), packet, callID, false, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showPlayedGame := types.NewPrimitiveBool(false)
	err = showPlayedGame.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read showPlayedGame from parameters. %s", err.Error()), packet, callID, false, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePreference(nil, packet, callID, publicMode, showGame, showPlayedGame)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
