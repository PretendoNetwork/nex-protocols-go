// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetlstFriendNotificationData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetlstFriendNotificationData == nil {
		globals.Logger.Warning("MatchmakeExtension::GetlstFriendNotificationData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	lstTypes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetlstFriendNotificationData(fmt.Errorf("Failed to read lstTypes from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetlstFriendNotificationData(nil, packet, callID, lstTypes)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
