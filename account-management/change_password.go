// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePassword sets the ChangePassword handler function
func (protocol *AccountManagementProtocol) ChangePassword(handler func(err error, client *nex.Client, callID uint32, strNewKey string)) {
	protocol.changePasswordHandler = handler
}

func (protocol *AccountManagementProtocol) handleChangePassword(packet nex.PacketInterface) {
	if protocol.changePasswordHandler == nil {
		globals.Logger.Warning("AccountManagement::ChangePassword not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strNewKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.changePasswordHandler(fmt.Errorf("Failed to read strNewKey from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.changePasswordHandler(nil, client, callID, strNewKey)
}
