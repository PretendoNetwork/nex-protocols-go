// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePlayedGames(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePlayedGames == nil {
		globals.Logger.Warning("Friends3DS::UpdatePlayedGames not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	playedGames := types.NewList[*friends_3ds_types.PlayedGame]()
	playedGames.Type = friends_3ds_types.NewPlayedGame()
	err = playedGames.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePlayedGames(fmt.Errorf("Failed to read playedGames from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePlayedGames(nil, packet, callID, playedGames)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
