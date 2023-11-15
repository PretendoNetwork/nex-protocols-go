// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSearchUnknownPlatformObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SearchUnknownPlatformObjects == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.SearchUnknownPlatformObjects(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
