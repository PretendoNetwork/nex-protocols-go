// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateCustomData sets the UpdateCustomData handler function
func (protocol *AccountManagementProtocol) UpdateCustomData(handler func(err error, client *nex.Client, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)) {
	protocol.updateCustomDataHandler = handler
}

func (protocol *AccountManagementProtocol) handleUpdateCustomData(packet nex.PacketInterface) {
	if protocol.updateCustomDataHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateCustomData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.updateCustomDataHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.updateCustomDataHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.updateCustomDataHandler(nil, client, callID, oPublicData, oPrivateData)
}
