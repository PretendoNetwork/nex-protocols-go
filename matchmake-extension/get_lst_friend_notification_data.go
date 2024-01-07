// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetlstFriendNotificationData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetlstFriendNotificationData == nil {
		globals.Logger.Warning("MatchmakeExtension::GetlstFriendNotificationData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstTypes := types.NewList[*types.PrimitiveU32]()
	lstTypes.Type = types.NewPrimitiveU32(0)
	err = lstTypes.ExtractFrom(parametersStream)
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
