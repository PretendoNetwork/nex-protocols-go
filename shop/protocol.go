// Package protocol implements the Shop protocol
package protocol

// * Stubbed, Kinnay documents this as being game-specific for Pokemon bank however Badge Arcade and Pokemon gen 7 uses this protocol as well
// TODO - Figure out more about this protocol, unsure if anything here is right

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the Protocol ID for the Shop protocol
	ProtocolID = 0xC8
)

// Protocol stores all the RMC method handlers for the Shop protocol and listens for requests
type Protocol struct {
	Server nex.ServerInterface
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Shop method ID: %#v\n", request.MethodID)
	}
}
