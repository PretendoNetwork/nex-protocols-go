// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdatePreference == nil {
		globals.Logger.Warning("Friends3DS::UpdatePreference not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	publicMode, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read publicMode from parameters. %s", err.Error()), packet, callID, false, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showGame, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), packet, callID, false, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showPlayedGame, err := parametersStream.ReadBool()
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
