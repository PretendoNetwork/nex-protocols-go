// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchUnknownPlatformObjects sets the SearchUnknownPlatformObjects handler function
func (protocol *Protocol) SearchUnknownPlatformObjects(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.searchUnknownPlatformObjectsHandler = handler
}

func (protocol *Protocol) handleSearchUnknownPlatformObjects(packet nex.PacketInterface) {
	if protocol.searchUnknownPlatformObjectsHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("DataStoreSuperMarioMaker::SearchUnknownPlatformObjects STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.searchUnknownPlatformObjectsHandler(nil, client, callID, packet.Payload())
}
