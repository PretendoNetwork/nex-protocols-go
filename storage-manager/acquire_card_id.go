// Package protocol implements the StorageManager protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAcquireCardID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AcquireCardID == nil {
		globals.Logger.Warning("StorageManager::AcquireCardID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.AcquireCardID(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
