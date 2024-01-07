// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.MarkFriendRequestsAsReceived == nil {
		globals.Logger.Warning("FriendsWiiU::MarkFriendRequestsAsReceived not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	ids := types.NewList[*types.PrimitiveU64]()
	ids.Type = types.NewPrimitiveU64(0)
	err = ids.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRequestBlockSettings(fmt.Errorf("Failed to read ids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.MarkFriendRequestsAsReceived(nil, packet, callID, ids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
