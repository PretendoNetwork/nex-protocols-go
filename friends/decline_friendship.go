// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeclineFriendship sets the DeclineFriendship handler function
func (protocol *Protocol) DeclineFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32) uint32) {
	protocol.declineFriendshipHandler = handler
}

func (protocol *Protocol) handleDeclineFriendship(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.declineFriendshipHandler == nil {
		globals.Logger.Warning("Friends::DeclineFriendship not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.declineFriendshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.declineFriendshipHandler(nil, packet, callID, uiPlayer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
