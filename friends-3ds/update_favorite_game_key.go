// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	if protocol.UpdateFavoriteGameKey == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::UpdateFavoriteGameKey not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gameKey := friends_3ds_types.NewGameKey()

	err := gameKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateFavoriteGameKey(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateFavoriteGameKey(nil, packet, callID, gameKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
