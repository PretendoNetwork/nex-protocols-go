// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrivateData sets the GetPrivateData handler function
func (protocol *Protocol) GetPrivateData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getPrivateDataHandler = handler
}

func (protocol *Protocol) handleGetPrivateData(packet nex.PacketInterface) {
	if protocol.getPrivateDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetPrivateData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getPrivateDataHandler(nil, client, callID)
}
