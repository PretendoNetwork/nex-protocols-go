// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMultiplePublicData sets the GetMultiplePublicData handler function
func (protocol *Protocol) GetMultiplePublicData(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []uint32) uint32) {
	protocol.getMultiplePublicDataHandler = handler
}

func (protocol *Protocol) handleGetMultiplePublicData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMultiplePublicDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetMultiplePublicData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipals, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getMultiplePublicDataHandler(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMultiplePublicDataHandler(nil, packet, callID, lstPrincipals)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
