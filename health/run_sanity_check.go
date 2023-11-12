// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RunSanityCheck sets the RunSanityCheck handler function
func (protocol *Protocol) RunSanityCheck(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.runSanityCheckHandler = handler
}

func (protocol *Protocol) handleRunSanityCheck(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.runSanityCheckHandler == nil {
		globals.Logger.Warning("Health::RunSanityCheck not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.runSanityCheckHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
