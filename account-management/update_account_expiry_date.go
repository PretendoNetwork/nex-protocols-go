// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountExpiryDate sets the UpdateAccountExpiryDate handler function
func (protocol *Protocol) UpdateAccountExpiryDate(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtExpiry *nex.DateTime, strExpiredMessage string) uint32) {
	protocol.updateAccountExpiryDateHandler = handler
}

func (protocol *Protocol) handleUpdateAccountExpiryDate(packet nex.PacketInterface) {
	if protocol.updateAccountExpiryDateHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountExpiryDate not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
