// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetName == nil {
		globals.Logger.Warning("AccountManagement::GetName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.GetName(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetName(nil, packet, callID, idPrincipal)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
