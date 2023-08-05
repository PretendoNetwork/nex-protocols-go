// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueIDWithPassword sets the AcquireNexUniqueIDWithPassword handler function
func (protocol *Protocol) AcquireNexUniqueIDWithPassword(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.acquireNexUniqueIDWithPasswordHandler = handler
}

func (protocol *Protocol) handleAcquireNexUniqueIDWithPassword(packet nex.PacketInterface) {
	if protocol.acquireNexUniqueIDWithPasswordHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueIDWithPassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.acquireNexUniqueIDWithPasswordHandler(nil, client, callID)
}
