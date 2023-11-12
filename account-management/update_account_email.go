// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountEmail sets the UpdateAccountEmail handler function
func (protocol *Protocol) UpdateAccountEmail(handler func(err error, packet nex.PacketInterface, callID uint32, strName string) uint32) {
	protocol.updateAccountEmailHandler = handler
}

func (protocol *Protocol) handleUpdateAccountEmail(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateAccountEmailHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountEmail not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateAccountEmailHandler(fmt.Errorf("Failed to read strName from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateAccountEmailHandler(nil, packet, callID, strName)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
