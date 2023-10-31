// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
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
	Server *nex.Server
	secureConnectionProtocol
	getMaintenanceStatusHandler func(err error, packet nex.PacketInterface, callID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.secureConnectionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetMaintenanceStatus:
		go protocol.handleGetMaintenanceStatus(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported SecureConnectionNintendoBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Secure Connection (Nintendo Badge Arcade) protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.secureConnectionProtocol.Server = server

	protocol.Setup()

	return protocol
}
