// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendNotificationData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetFriendNotificationData == nil {
		globals.Logger.Warning("MatchmakeExtension::GetFriendNotificationData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType, err := parametersStream.ReadInt32LE()
	if err != nil {
		errorCode = protocol.GetFriendNotificationData(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetFriendNotificationData(nil, packet, callID, uiType)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
