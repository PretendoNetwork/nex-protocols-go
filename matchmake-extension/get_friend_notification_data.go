// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendNotificationData sets the GetFriendNotificationData handler function
func (protocol *Protocol) GetFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType int32) uint32) {
	protocol.getFriendNotificationDataHandler = handler
}

func (protocol *Protocol) handleGetFriendNotificationData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetFriendNotificationData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType, err := parametersStream.ReadInt32LE()
	if err != nil {
		errorCode = protocol.getFriendNotificationDataHandler(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendNotificationDataHandler(nil, packet, callID, uiType)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
