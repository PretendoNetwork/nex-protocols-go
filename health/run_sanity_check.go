// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RunSanityCheck sets the RunSanityCheck handler function
func (protocol *Protocol) RunSanityCheck(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.RunSanityCheckHandler = handler
}

func (protocol *Protocol) handleRunSanityCheck(packet nex.PacketInterface) {
	if protocol.RunSanityCheckHandler == nil {
		globals.Logger.Warning("Health::RunSanityCheck not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.RunSanityCheckHandler(nil, client, callID)
}
