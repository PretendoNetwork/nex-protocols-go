// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsActiveGame sets the IsActiveGame handler function
func (protocol *Protocol) IsActiveGame(handler func(err error, packet nex.PacketInterface, callID uint32, pids []uint32, gameKey *friends_3ds_types.GameKey) uint32) {
	protocol.isActiveGameHandler = handler
}

func (protocol *Protocol) handleIsActiveGame(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.isActiveGameHandler == nil {
		globals.Logger.Warning("Friends3DS::IsActiveGame not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.isActiveGameHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	gameKey, err := parametersStream.ReadStructure(friends_3ds_types.NewGameKey())
	if err != nil {
		errorCode = protocol.isActiveGameHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.isActiveGameHandler(nil, packet, callID, pids, gameKey.(*friends_3ds_types.GameKey))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
