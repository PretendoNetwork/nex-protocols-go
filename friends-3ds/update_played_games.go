// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePlayedGames sets the UpdatePlayedGames handler function
func (protocol *Protocol) UpdatePlayedGames(handler func(err error, packet nex.PacketInterface, callID uint32, playedGames []*friends_3ds_types.PlayedGame) uint32) {
	protocol.updatePlayedGamesHandler = handler
}

func (protocol *Protocol) handleUpdatePlayedGames(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updatePlayedGamesHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePlayedGames not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	playedGames, err := parametersStream.ReadListStructure(friends_3ds_types.NewPlayedGame())
	if err != nil {
		errorCode = protocol.updatePlayedGamesHandler(fmt.Errorf("Failed to read playedGames from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updatePlayedGamesHandler(nil, packet, callID, playedGames.([]*friends_3ds_types.PlayedGame))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
