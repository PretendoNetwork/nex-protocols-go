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
	if protocol.UpdatePlayedGames == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::UpdatePlayedGames not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	playedGames := types.NewList[*friends_3ds_types.PlayedGame]()
	playedGames.Type = friends_3ds_types.NewPlayedGame()

	err := playedGames.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePlayedGames(fmt.Errorf("Failed to read playedGames from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdatePlayedGames(nil, packet, callID, playedGames)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
