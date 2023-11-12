// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateCustomData sets the UpdateCustomData handler function
func (protocol *Protocol) UpdateCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32) {
	protocol.updateCustomDataHandler = handler
}

func (protocol *Protocol) handleUpdateCustomData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateCustomDataHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateCustomData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateCustomDataHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateCustomDataHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateCustomDataHandler(nil, packet, callID, oPublicData, oPrivateData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
