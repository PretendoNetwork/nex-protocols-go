// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SimpleFindByID sets the SimpleFindByID handler function
func (protocol *Protocol) SimpleFindByID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.simpleFindByIDHandler = handler
}

func (protocol *Protocol) handleSimpleFindByID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.simpleFindByIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SimpleFindByID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SimpleFindByID STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.simpleFindByIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
