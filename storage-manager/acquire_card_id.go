// Package protocol implements the StorageManager protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireCardID sets the AcquireCardID handler function
func (protocol *Protocol) AcquireCardID(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.acquireCardIDHandler = handler
}

func (protocol *Protocol) handleAcquireCardID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acquireCardIDHandler == nil {
		globals.Logger.Warning("StorageManager::AcquireCardID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.acquireCardIDHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
