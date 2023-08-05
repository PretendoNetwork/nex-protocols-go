// Package protocol implements the DataStoreSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSharedDataMeta sets the GetSharedDataMeta handler function
func (protocol *Protocol) GetSharedDataMeta(handler func(err error, client *nex.Client, callID uint32, pakcetPayload []byte) uint32) {
	protocol.getSharedDataMetaHandler = handler
}

func (protocol *Protocol) handleGetSharedDataMeta(packet nex.PacketInterface) {
	if protocol.getSharedDataMetaHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetSharedDataMeta not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("DataStoreSuperSmashBros4::GetSharedDataMeta STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.getSharedDataMetaHandler(nil, client, callID, packet.Payload())
}
