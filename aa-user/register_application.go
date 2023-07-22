// Package aauser implements the AAUser NEX protocol
package aauser

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterApplication sets the RegisterApplication handler function
func (protocol *AAUserProtocol) RegisterApplication(handler func(err error, client *nex.Client, callID uint32, titleID uint64)) {
	protocol.registerApplicationHandler = handler
}

func (protocol *AAUserProtocol) handleRegisterApplication(packet nex.PacketInterface) {
	if protocol.registerApplicationHandler == nil {
		globals.Logger.Warning("AAUser::RegisterApplication not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	titleID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.registerApplicationHandler(fmt.Errorf("Failed to read titleID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.registerApplicationHandler(nil, client, callID, titleID)
}
