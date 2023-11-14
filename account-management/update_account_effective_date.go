// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccountEffectiveDate sets the UpdateAccountEffectiveDate handler function
func (protocol *Protocol) UpdateAccountEffectiveDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) uint32) {
	protocol.updateAccountEffectiveDateHandler = handler
}

func (protocol *Protocol) handleUpdateAccountEffectiveDate(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateAccountEffectiveDateHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountEffectiveDate not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dtEffectiveFrom, err := parametersStream.ReadDateTime()
	if err != nil {
		errorCode = protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read dtEffectiveFrom from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strNotEffectiveMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateAccountEffectiveDateHandler(fmt.Errorf("Failed to read strNotEffectiveMessage from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateAccountEffectiveDateHandler(nil, packet, callID, idPrincipal, dtEffectiveFrom, strNotEffectiveMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
