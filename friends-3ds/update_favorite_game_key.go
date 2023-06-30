// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateFavoriteGameKey sets the UpdateFavoriteGameKey handler function
func (protocol *Friends3DSProtocol) UpdateFavoriteGameKey(handler func(err error, client *nex.Client, callID uint32, gameKey *friends_3ds_types.GameKey)) {
	protocol.UpdateFavoriteGameKeyHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	if protocol.UpdateFavoriteGameKeyHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateFavoriteGameKey not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gameKey, err := parametersStream.ReadStructure(friends_3ds_types.NewGameKey())
	if err != nil {
		go protocol.UpdateFavoriteGameKeyHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateFavoriteGameKeyHandler(nil, client, callID, gameKey.(*friends_3ds_types.GameKey))
}
