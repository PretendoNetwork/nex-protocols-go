// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePasswordByGuest sets the ChangePasswordByGuest handler function
func (protocol *Protocol) ChangePasswordByGuest(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) uint32) {
	protocol.changePasswordByGuestHandler = handler
}

func (protocol *Protocol) handleChangePasswordByGuest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.changePasswordByGuestHandler == nil {
		globals.Logger.Warning("AccountManagement::ChangePasswordByGuest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, "", "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, "", "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.changePasswordByGuestHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, "", "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.changePasswordByGuestHandler(nil, packet, callID, strPrincipalName, strKey, strEmail)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
