// Package protocol implements the Monitoring protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Monitoring protocol
	ProtocolID = 0x13

	// MethodPingDaemon is the method ID for the method PingDaemon
	MethodPingDaemon = 0x1

	// MethodGetClusterMembers is the method ID for the method GetClusterMembers
	MethodGetClusterMembers = 0x2
)

// Protocol handles the Monitoring protocol
type Protocol struct {
	Server            nex.ServerInterface
	PingDaemon        func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetClusterMembers func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
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
	case MethodPingDaemon:
		go protocol.handlePingDaemon(packet)
	case MethodGetClusterMembers:
		go protocol.handleGetClusterMembers(packet)
	default:
		fmt.Printf("Unsupported Monitoring method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Monitoring protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
