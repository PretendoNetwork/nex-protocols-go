// Package protocol implements the Remote Log Device protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the RemoteLogDevice protocol
	ProtocolID = 0x1

	// MethodLog is the method ID for the method Log
	MethodLog = 0x1
)

// Protocol handles the RemoteLogDevice protocol
type Protocol struct {
	Server     *nex.Server
	logHandler func(err error, packet nex.PacketInterface, callID uint32, strLine string) uint32
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
	case MethodLog:
		go protocol.handleLog(packet)
	default:
		fmt.Printf("Unsupported RemoteLogDevice method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Remote Log Device protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
