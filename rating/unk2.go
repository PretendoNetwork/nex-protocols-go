// Package protocol implements the Rating protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// TODO: find name if possible
// TODO: implement correctly
func (protocol *Protocol) handleUnk2(packet nex.PacketInterface) {
	if protocol.Unk2 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Rating::Unk2 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.Unk2(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
