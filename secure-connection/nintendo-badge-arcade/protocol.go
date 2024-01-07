// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection (Nintendo Badge Arcade) protocol. ID is the same as the Secure Connection Protocol
	ProtocolID = 0xB

	// MethodGetMaintenanceStatus is the method ID for GetMaintenanceStatus
	MethodGetMaintenanceStatus = 0x9
)

var patchedMethods = []uint32{
	MethodGetMaintenanceStatus,
}

type secureConnectionProtocol = secure_connection.Protocol

// Protocol stores all the RMC method handlers for the Secure Connection (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the SecureConnection Protocol
type Protocol struct {
	server nex.ServerInterface
	secureConnectionProtocol
	GetMaintenanceStatus func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, message.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.secureConnectionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodGetMaintenanceStatus:
		protocol.handleGetMaintenanceStatus(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported SecureConnectionNintendoBadgeArcade method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Secure Connection (Nintendo Badge Arcade) protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.secureConnectionProtocol.SetServer(server)

	protocol.Setup()

	return protocol
}
