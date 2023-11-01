// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateFavoriteGameKey sets the UpdateFavoriteGameKey handler function
func (protocol *Protocol) UpdateFavoriteGameKey(handler func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) uint32) {
	protocol.updateFavoriteGameKeyHandler = handler
}

func (protocol *Protocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateFavoriteGameKeyHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateFavoriteGameKey not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gameKey, err := parametersStream.ReadStructure(friends_3ds_types.NewGameKey())
	if err != nil {
		errorCode = protocol.updateFavoriteGameKeyHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateFavoriteGameKeyHandler(nil, packet, callID, gameKey.(*friends_3ds_types.GameKey))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
