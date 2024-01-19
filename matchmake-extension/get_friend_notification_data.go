// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendNotificationData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetFriendNotificationData == nil {
		globals.Logger.Warning("MatchmakeExtension::GetFriendNotificationData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiType := types.NewPrimitiveS32(0)
	err = uiType.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetFriendNotificationData(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetFriendNotificationData(nil, packet, callID, uiType)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
