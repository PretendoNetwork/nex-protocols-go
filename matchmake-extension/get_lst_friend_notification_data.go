// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetlstFriendNotificationData(packet nex.PacketInterface) {
	if protocol.GetlstFriendNotificationData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::GetlstFriendNotificationData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstTypes := types.NewList[*types.PrimitiveU32]()
	lstTypes.Type = types.NewPrimitiveU32(0)

	err := lstTypes.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetlstFriendNotificationData(fmt.Errorf("Failed to read lstTypes from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetlstFriendNotificationData(nil, packet, callID, lstTypes)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
