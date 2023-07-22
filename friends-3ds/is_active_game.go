// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsActiveGame sets the IsActiveGame handler function
func (protocol *Friends3DSProtocol) IsActiveGame(handler func(err error, client *nex.Client, callID uint32, pids []uint32, gameKey *friends_3ds_types.GameKey)) {
	protocol.isActiveGameHandler = handler
}

func (protocol *Friends3DSProtocol) handleIsActiveGame(packet nex.PacketInterface) {
	if protocol.isActiveGameHandler == nil {
		globals.Logger.Warning("Friends3DS::IsActiveGame not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.isActiveGameHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	gameKey, err := parametersStream.ReadStructure(friends_3ds_types.NewGameKey())
	if err != nil {
		go protocol.isActiveGameHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.isActiveGameHandler(nil, client, callID, pids, gameKey.(*friends_3ds_types.GameKey))
}
