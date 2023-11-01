// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchUnknownPlatformObjects sets the SearchUnknownPlatformObjects handler function
func (protocol *Protocol) SearchUnknownPlatformObjects(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.searchUnknownPlatformObjectsHandler = handler
}

func (protocol *Protocol) handleSearchUnknownPlatformObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchUnknownPlatformObjectsHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.searchUnknownPlatformObjectsHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
