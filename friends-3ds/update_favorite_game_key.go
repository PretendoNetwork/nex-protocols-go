// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateFavoriteGameKey == nil {
		globals.Logger.Warning("Friends3DS::UpdateFavoriteGameKey not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gameKey := friends_3ds_types.NewGameKey()
	err = gameKey.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateFavoriteGameKey(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateFavoriteGameKey(nil, packet, callID, gameKey)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
