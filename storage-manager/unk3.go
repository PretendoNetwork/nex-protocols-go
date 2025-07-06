// Package protocol implements the StorageManager protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// TODO - Find name if possible
// TODO - Implement correctly
func (protocol *Protocol) handleUnk3(packet nex.PacketInterface) {
	if protocol.Unk3 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "StorageManager::Unk3 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.Unk3(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
