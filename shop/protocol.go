package shop

// * Stubbed, Kinnay documents this as being game-specific for Pokemon bank however Badge Arcade and Pokemon gen 7 uses this protocol as well
// TODO - Figure out more about this protocol, unsure if anything here is right

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the Protocol ID for the Shop (Badge Arcade) protocol
	ProtocolID = 0xC8
)

// ShopProtocol handles the Shop (Badge Arcade) nex protocol
type ShopProtocol struct {
	Server *nex.Server
}

// Setup initializes the protocol
func (protocol *ShopProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *ShopProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Shop method ID: %#v\n", request.MethodID())
	}
}
