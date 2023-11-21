// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPrincipalIDByLocalFriendCode(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetPrincipalIDByLocalFriendCode == nil {
		globals.Logger.Warning("Friends3DS::GetPrincipalIDByLocalFriendCode not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetPrincipalIDByLocalFriendCode(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetPrincipalIDByLocalFriendCode(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPrincipalIDByLocalFriendCode(nil, packet, callID, lfc, lfcList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
