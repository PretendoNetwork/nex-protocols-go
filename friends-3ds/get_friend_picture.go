// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPicture sets the GetFriendPicture handler function
func (protocol *Protocol) GetFriendPicture(handler func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) uint32) {
	protocol.getFriendPictureHandler = handler
}

func (protocol *Protocol) handleGetFriendPicture(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendPictureHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPicture not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getFriendPictureHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendPictureHandler(nil, packet, callID, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
