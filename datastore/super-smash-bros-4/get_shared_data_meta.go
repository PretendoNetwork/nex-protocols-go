// Package protocol implements the DataStoreSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSharedDataMeta(packet nex.PacketInterface) {
	if protocol.GetSharedDataMeta == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetSharedDataMeta not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("DataStoreSuperSmashBros4::GetSharedDataMeta STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	rmcMessage, errorCode := protocol.GetSharedDataMeta(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
