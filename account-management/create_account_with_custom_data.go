// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateAccountWithCustomData sets the CreateAccountWithCustomData handler function
func (protocol *Protocol) CreateAccountWithCustomData(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)) {
	protocol.createAccountWithCustomDataHandler = handler
}

func (protocol *Protocol) handleCreateAccountWithCustomData(packet nex.PacketInterface) {
	if protocol.createAccountWithCustomDataHandler == nil {
		globals.Logger.Warning("AccountManagement::CreateAccountWithCustomData not implemented")
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
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil, nil)
		return
	}

	go protocol.createAccountWithCustomDataHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail, oPublicData, oPrivateData)
}
