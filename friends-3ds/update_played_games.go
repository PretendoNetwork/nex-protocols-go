// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePlayedGames sets the UpdatePlayedGames handler function
func (protocol *Friends3DSProtocol) UpdatePlayedGames(handler func(err error, client *nex.Client, callID uint32, playedGames []*friends_3ds_types.PlayedGame)) {
	protocol.updatePlayedGamesHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdatePlayedGames(packet nex.PacketInterface) {
	if protocol.updatePlayedGamesHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePlayedGames not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	playedGames, err := parametersStream.ReadListStructure(friends_3ds_types.NewPlayedGame())
	if err != nil {
		go protocol.updatePlayedGamesHandler(fmt.Errorf("Failed to read playedGames from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updatePlayedGamesHandler(nil, client, callID, playedGames.([]*friends_3ds_types.PlayedGame))
}
