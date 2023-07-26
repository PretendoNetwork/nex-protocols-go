// Package protocol implements the StorageManager protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireCardID sets the AcquireCardID handler function
func (protocol *Protocol) AcquireCardID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.acquireCardIDHandler = handler
}

func (protocol *Protocol) handleAcquireCardID(packet nex.PacketInterface) {
	if protocol.acquireCardIDHandler == nil {
		globals.Logger.Warning("StorageManager::AcquireCardID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.acquireCardIDHandler(nil, client, callID)
}
