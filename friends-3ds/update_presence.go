// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdatePresence == nil {
		globals.Logger.Warning("Friends3DS::UpdatePresence not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresence, err := parametersStream.ReadStructure(friends_3ds_types.NewNintendoPresence())
	if err != nil {
		_, errorCode = protocol.UpdatePresence(fmt.Errorf("Failed to read nintendoPresence from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showGame, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.UpdatePresence(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePresence(nil, packet, callID, nintendoPresence.(*friends_3ds_types.NintendoPresence), showGame)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
