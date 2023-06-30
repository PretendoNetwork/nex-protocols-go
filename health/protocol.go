// Package health implements the Health NEX protocol
package health

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

// HealthProtocol handles the Health protocol
type HealthProtocol struct {
	Server                 *nex.Server
	PingDaemonHandler      func(err error, client *nex.Client, callID uint32)
	PingDatabaseHandler    func(err error, client *nex.Client, callID uint32)
	RunSanityCheckHandler  func(err error, client *nex.Client, callID uint32)
	FixSanityErrorsHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *HealthProtocol) Setup() {
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

// NewHealthProtocol returns a new HealthProtocol
func NewHealthProtocol(server *nex.Server) *HealthProtocol {
	protocol := &HealthProtocol{Server: server}

	protocol.Setup()

	return protocol
}
