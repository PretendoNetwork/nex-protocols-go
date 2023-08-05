// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateFavoriteGameKey sets the UpdateFavoriteGameKey handler function
func (protocol *Protocol) UpdateFavoriteGameKey(handler func(err error, client *nex.Client, callID uint32, gameKey *friends_3ds_types.GameKey) uint32) {
	protocol.updateFavoriteGameKeyHandler = handler
}

func (protocol *Protocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	if protocol.updateFavoriteGameKeyHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateFavoriteGameKey not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gameKey, err := parametersStream.ReadStructure(friends_3ds_types.NewGameKey())
	if err != nil {
		go protocol.updateFavoriteGameKeyHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateFavoriteGameKeyHandler(nil, client, callID, gameKey.(*friends_3ds_types.GameKey))
}
