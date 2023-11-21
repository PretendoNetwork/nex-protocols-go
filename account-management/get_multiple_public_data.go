// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetMultiplePublicData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetMultiplePublicData == nil {
		globals.Logger.Warning("AccountManagement::GetMultiplePublicData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipals, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetMultiplePublicData(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetMultiplePublicData(nil, packet, callID, lstPrincipals)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
