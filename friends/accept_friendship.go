// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAcceptFriendship(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AcceptFriendship == nil {
		globals.Logger.Warning("Friends::AcceptFriendship not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.AcceptFriendship(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AcceptFriendship(nil, packet, callID, uiPlayer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
