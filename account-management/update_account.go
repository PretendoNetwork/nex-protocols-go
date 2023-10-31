// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAccount sets the UpdateAccount handler function
func (protocol *Protocol) UpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32) {
	protocol.updateAccountHandler = handler
}

func (protocol *Protocol) handleUpdateAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, "", "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, "", "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateAccountHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, "", "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateAccountHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, "", "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateAccountHandler(nil, packet, callID, strKey, strEmail, oPublicData, oPrivateData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
