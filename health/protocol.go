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
	Server          nex.ServerInterface
	PingDaemon      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	PingDatabase    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	RunSanityCheck  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	FixSanityErrors func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
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
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
