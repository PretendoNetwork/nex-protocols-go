// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountName sets the UpdateAccountName handler function
func (protocol *Protocol) UpdateAccountName(handler func(err error, client *nex.Client, callID uint32, strName string) uint32) {
	protocol.updateAccountNameHandler = handler
}

func (protocol *Protocol) handleUpdateAccountName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateAccountNameHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateAccountNameHandler(fmt.Errorf("Failed to read strName from parameters. %s", err.Error()), client, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateAccountNameHandler(nil, client, callID, strName)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
