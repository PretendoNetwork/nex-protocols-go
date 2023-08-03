// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SimpleFindByID sets the SimpleFindByID handler function
func (protocol *Protocol) SimpleFindByID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.simpleFindByIDHandler = handler
}

func (protocol *Protocol) handleSimpleFindByID(packet nex.PacketInterface) {
	if protocol.simpleFindByIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SimpleFindByID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SimpleFindByID STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.simpleFindByIDHandler(nil, client, callID, packet.Payload())
}
