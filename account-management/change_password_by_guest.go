// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePasswordByGuest sets the ChangePasswordByGuest handler function
func (protocol *AccountManagementProtocol) ChangePasswordByGuest(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, strEmail string)) {
	protocol.changePasswordByGuestHandler = handler
}

func (protocol *AccountManagementProtocol) handleChangePasswordByGuest(packet nex.PacketInterface) {
	if protocol.changePasswordByGuestHandler == nil {
		globals.Logger.Warning("AccountManagement::ChangePasswordByGuest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", "")
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", "")
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", "")
		return
	}

	go protocol.changePasswordByGuestHandler(nil, client, callID, strPrincipalName, strKey, strEmail)
}
