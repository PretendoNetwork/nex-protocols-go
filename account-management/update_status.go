// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateStatus(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateStatus == nil {
		globals.Logger.Warning("AccountManagement::UpdateStatus not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strStatus, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.UpdateStatus(fmt.Errorf("Failed to read strStatus from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateStatus(nil, packet, callID, strStatus)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
