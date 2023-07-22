// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountEffectiveDate sets the UpdateAccountEffectiveDate handler function
func (protocol *AccountManagementProtocol) UpdateAccountEffectiveDate(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string)) {
	protocol.updateAccountEffectiveDateHandler = handler
}

func (protocol *AccountManagementProtocol) handleUpdateAccountEffectiveDate(packet nex.PacketInterface) {
	if protocol.updateAccountEffectiveDateHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountEffectiveDate not implemented")
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
		go protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	dtEffectiveFrom, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read dtEffectiveFrom from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	strNotEffectiveMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read strNotEffectiveMessage from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	go protocol.updateAccountEffectiveDateHandler(nil, client, callID, idPrincipal, dtEffectiveFrom, strNotEffectiveMessage)
}
