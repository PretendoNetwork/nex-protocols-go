// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountName sets the UpdateAccountName handler function
func (protocol *AccountManagementProtocol) UpdateAccountName(handler func(err error, client *nex.Client, callID uint32, strName string)) {
	protocol.updateAccountNameHandler = handler
}

func (protocol *AccountManagementProtocol) handleUpdateAccountName(packet nex.PacketInterface) {
	if protocol.updateAccountNameHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountName not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateAccountNameHandler(fmt.Errorf("Failed to read strName from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.updateAccountNameHandler(nil, client, callID, strName)
}
