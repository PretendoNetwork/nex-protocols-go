// Package protocol implements the Health protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
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
	Server                 *nex.Server
	pingDaemonHandler      func(err error, packet nex.PacketInterface, callID uint32) uint32
	pingDatabaseHandler    func(err error, packet nex.PacketInterface, callID uint32) uint32
	runSanityCheckHandler  func(err error, packet nex.PacketInterface, callID uint32) uint32
	fixSanityErrorsHandler func(err error, packet nex.PacketInterface, callID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodPingDaemon:
				go protocol.handlePingDaemon(packet)
			case MethodPingDatabase:
				go protocol.handlePingDatabase(packet)
			case MethodRunSanityCheck:
				go protocol.handleRunSanityCheck(packet)
			case MethodFixSanityErrors:
				go protocol.handleFixSanityErrors(packet)
			default:
				fmt.Printf("Unsupported Health method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Health protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
