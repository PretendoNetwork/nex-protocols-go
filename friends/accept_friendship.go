// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptFriendship sets the AcceptFriendship handler function
func (protocol *Protocol) AcceptFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32) uint32) {
	protocol.acceptFriendshipHandler = handler
}

func (protocol *Protocol) handleAcceptFriendship(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acceptFriendshipHandler == nil {
		globals.Logger.Warning("Friends::AcceptFriendship not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.acceptFriendshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.acceptFriendshipHandler(nil, packet, callID, uiPlayer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
