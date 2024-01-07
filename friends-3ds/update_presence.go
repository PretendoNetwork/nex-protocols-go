// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePresence == nil {
		globals.Logger.Warning("Friends3DS::UpdatePresence not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	nintendoPresence := friends_3ds_types.NewNintendoPresence()
	err = nintendoPresence.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePresence(fmt.Errorf("Failed to read nintendoPresence from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showGame := types.NewPrimitiveBool(false)
	err = showGame.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePresence(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePresence(nil, packet, callID, nintendoPresence, showGame)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
