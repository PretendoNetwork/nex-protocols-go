// Package protocol implements the Health protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Health protocol
	ProtocolID = 0x12

	// MethodPingDaemon is the method ID for the method PingDaemon
	MethodPingDaemon = 0x1

	// MethodPingDatabase is the method ID for the method PingDatabase
	MethodPingDatabase = 0x2

	// MethodRunSanityCheck is the method ID for the method RunSanityCheck
	MethodRunSanityCheck = 0x3

	// MethodFixSanityErrors is the method ID for the method FixSanityErrors
	MethodFixSanityErrors = 0x4
)

// Protocol handles the Health protocol
type Protocol struct {
	server          nex.ServerInterface
	PingDaemon      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	PingDatabase    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	RunSanityCheck  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	FixSanityErrors func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Health protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerPingDatabase(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerRunSanityCheck(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerFixSanityErrors(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerPingDaemon sets the handler for the PingDaemon method
func (protocol *Protocol) SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.PingDaemon = handler
}

// SetHandlerPingDatabase sets the handler for the PingDatabase method
func (protocol *Protocol) SetHandlerPingDatabase(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.PingDatabase = handler
}

// SetHandlerRunSanityCheck sets the handler for the RunSanityCheck method
func (protocol *Protocol) SetHandlerRunSanityCheck(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.RunSanityCheck = handler
}

// SetHandlerFixSanityErrors sets the handler for the FixSanityErrors method
func (protocol *Protocol) SetHandlerFixSanityErrors(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.FixSanityErrors = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodPingDaemon:
				protocol.handlePingDaemon(packet)
			case MethodPingDatabase:
				protocol.handlePingDatabase(packet)
			case MethodRunSanityCheck:
				protocol.handleRunSanityCheck(packet)
			case MethodFixSanityErrors:
				protocol.handleFixSanityErrors(packet)
			default:
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Health method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Health protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
