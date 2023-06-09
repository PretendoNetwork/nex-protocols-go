package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateFavoriteGameKey sets the UpdateFavoriteGameKey handler function
func (protocol *Friends3DSProtocol) UpdateFavoriteGameKey(handler func(err error, client *nex.Client, callID uint32, gameKey *GameKey)) {
	protocol.UpdateFavoriteGameKeyHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdateFavoriteGameKey(packet nex.PacketInterface) {
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

	gameKey, err := parametersStream.ReadStructure(NewGameKey())
	if err != nil {
		go protocol.UpdateFavoriteGameKeyHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateFavoriteGameKeyHandler(nil, client, callID, gameKey.(*GameKey))
}
