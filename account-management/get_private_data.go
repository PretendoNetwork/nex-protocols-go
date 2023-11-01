// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrivateData sets the GetPrivateData handler function
func (protocol *Protocol) GetPrivateData(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getPrivateDataHandler = handler
}

func (protocol *Protocol) handleGetPrivateData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPrivateDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetPrivateData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getPrivateDataHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
