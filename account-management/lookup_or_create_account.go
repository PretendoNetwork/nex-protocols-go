// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LookupOrCreateAccount sets the LookupOrCreateAccount handler function
func (protocol *AccountManagementProtocol) LookupOrCreateAccount(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)) {
	protocol.lookupOrCreateAccountHandler = handler
}

func (protocol *AccountManagementProtocol) handleLookupOrCreateAccount(packet nex.PacketInterface) {
	if protocol.lookupOrCreateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::LookupOrCreateAccount not implemented")
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
		go protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	oAuthData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.lookupOrCreateAccountHandler(fmt.Errorf("Failed to read oAuthData from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	go protocol.lookupOrCreateAccountHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail, oAuthData)
}
