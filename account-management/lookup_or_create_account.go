// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LookupOrCreateAccount sets the LookupOrCreateAccount handler function
func (protocol *Protocol) LookupOrCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) uint32) {
	protocol.lookupOrCreateAccountHandler = handler
}

func (protocol *Protocol) handleLookupOrCreateAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.lookupOrCreateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::LookupOrCreateAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oAuthData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read oAuthData from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.lookupOrCreateAccountHandler(nil, packet, callID, strPrincipalName, strKey, uiGroups, strEmail, oAuthData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
