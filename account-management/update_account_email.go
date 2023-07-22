// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountEmail sets the UpdateAccountEmail handler function
func (protocol *AccountManagementProtocol) UpdateAccountEmail(handler func(err error, client *nex.Client, callID uint32, strName string)) {
	protocol.updateAccountEmailHandler = handler
}

func (protocol *AccountManagementProtocol) handleUpdateAccountEmail(packet nex.PacketInterface) {
	if protocol.updateAccountEmailHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountEmail not implemented")
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
		go protocol.updateAccountEmailHandler(fmt.Errorf("Failed to read strName from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.updateAccountEmailHandler(nil, client, callID, strName)
}
