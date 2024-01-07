// Package protocol implements the Service Item protocol
package protocol

// TODO - Figure out more about this protocol, unsure if anything here is right

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the Protocol ID for the Service Item protocol
	ProtocolID = 0x77
)

// Protocol stores all the RMC method handlers for the Service Item protocol and listens for requests
type Protocol struct {
	server nex.ServerInterface
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported ServiceItem method ID: %#v\n", request.MethodID)
	}
}
