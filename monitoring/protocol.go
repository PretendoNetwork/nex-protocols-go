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
	Server                   *nex.Server
	pingDaemonHandler        func(err error, client *nex.Client, callID uint32)
	getClusterMembersHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPingDaemon:
		go protocol.handlePingDaemon(packet)
	case MethodGetClusterMembers:
		go protocol.handleGetClusterMembers(packet)
	default:
		fmt.Printf("Unsupported Monitoring method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Monitoring protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
