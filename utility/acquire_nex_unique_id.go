// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueID sets the AcquireNexUniqueID handler function
func (protocol *Protocol) AcquireNexUniqueID(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.acquireNexUniqueIDHandler = handler
}

func (protocol *Protocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	if protocol.acquireNexUniqueIDHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.acquireNexUniqueIDHandler(nil, client, callID)
}
