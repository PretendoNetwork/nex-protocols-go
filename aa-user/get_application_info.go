// Package aauser implements the AAUser NEX protocol
package aauser

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationInfo sets the GetApplicationInfo handler function
func (protocol *AAUserProtocol) GetApplicationInfo(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getApplicationInfoHandler = handler
}

func (protocol *AAUserProtocol) handleGetApplicationInfo(packet nex.PacketInterface) {
	if protocol.getApplicationInfoHandler == nil {
		globals.Logger.Warning("AAUser::GetApplicationInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getApplicationInfoHandler(nil, client, callID)
}
