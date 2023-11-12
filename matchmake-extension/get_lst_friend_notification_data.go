// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetlstFriendNotificationData sets the GetlstFriendNotificationData handler function
func (protocol *Protocol) GetlstFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, lstTypes []uint32) uint32) {
	protocol.getlstFriendNotificationDataHandler = handler
}

func (protocol *Protocol) handleGetlstFriendNotificationData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getlstFriendNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetlstFriendNotificationData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstTypes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getlstFriendNotificationDataHandler(fmt.Errorf("Failed to read lstTypes from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getlstFriendNotificationDataHandler(nil, packet, callID, lstTypes)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
