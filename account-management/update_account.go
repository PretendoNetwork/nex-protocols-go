// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccount sets the UpdateAccount handler function
func (protocol *Protocol) UpdateAccount(handler func(err error, client *nex.Client, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)) {
	protocol.updateAccountHandler = handler
}

func (protocol *Protocol) handleUpdateAccount(packet nex.PacketInterface) {
	if protocol.updateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccount not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", nil, nil)
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", nil, nil)
		return
	}

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.updateAccountHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), client, callID, "", "", nil, nil)
		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.updateAccountHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), client, callID, "", "", nil, nil)
		return
	}

	go protocol.updateAccountHandler(nil, client, callID, strKey, strEmail, oPublicData, oPrivateData)
}
