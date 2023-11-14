// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendBylstPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddFriendBylstPrincipalID == nil {
		globals.Logger.Warning("Friends3DS::AddFriendBylstPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.AddFriendBylstPrincipalID(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pids, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.AddFriendBylstPrincipalID(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddFriendBylstPrincipalID(nil, packet, callID, lfc, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
