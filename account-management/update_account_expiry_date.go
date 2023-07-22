// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountExpiryDate sets the UpdateAccountExpiryDate handler function
func (protocol *AccountManagementProtocol) UpdateAccountExpiryDate(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtExpiry *nex.DateTime, strExpiredMessage string)) {
	protocol.updateAccountExpiryDateHandler = handler
}

func (protocol *AccountManagementProtocol) handleUpdateAccountExpiryDate(packet nex.PacketInterface) {
	if protocol.updateAccountExpiryDateHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountExpiryDate not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateAccountExpiryDateHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	dtExpiry, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.updateAccountExpiryDateHandler(fmt.Errorf("Failed to read dtExpiry from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	strExpiredMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateAccountExpiryDateHandler(fmt.Errorf("Failed to read strExpiredMessage from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	go protocol.updateAccountExpiryDateHandler(nil, client, callID, idPrincipal, dtExpiry, strExpiredMessage)
}
