// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleChangePassword(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ChangePassword == nil {
		globals.Logger.Warning("AccountManagement::ChangePassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strNewKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.ChangePassword(fmt.Errorf("Failed to read strNewKey from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.ChangePassword(nil, packet, callID, strNewKey)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
