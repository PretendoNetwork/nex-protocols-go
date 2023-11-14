// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendUserStatuses(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetFriendUserStatuses == nil {
		globals.Logger.Warning("Subscriber::GetFriendUserStatuses not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt8()
	if err != nil {
		errorCode = protocol.GetFriendUserStatuses(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetFriendUserStatuses(nil, packet, callID, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
