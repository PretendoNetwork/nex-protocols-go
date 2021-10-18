package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// HealthProtocolID is the protocol ID for the Health protocol
	HealthProtocolID = 0x12

	// HealthMethodPingDaemon is the method ID for the method PingDaemon
	HealthMethodPingDaemon = 0x1

	// HealthMethodPingDatabase is the method ID for the method PingDatabase
	HealthMethodPingDatabase = 0x2

	// HealthMethodRunSanityCheck is the method ID for the method RunSanityCheck
	HealthMethodRunSanityCheck = 0x3

	// HealthMethodFixSanityErrors is the method ID for the method FixSanityErrors
	HealthMethodFixSanityErrors = 0x4
)

// HealthProtocol handles the Health protocol
type HealthProtocol struct {
	server                 *nex.Server
	PingDaemonHandler      func(err error, client *nex.Client, callID uint32)
	PingDatabaseHandler    func(err error, client *nex.Client, callID uint32)
	RunSanityCheckHandler  func(err error, client *nex.Client, callID uint32)
	FixSanityErrorsHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (healthProtocol *HealthProtocol) Setup() {
	nexServer := healthProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if HealthProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case HealthMethodPingDaemon:
				go healthProtocol.handlePingDaemon(packet)
				break
			case HealthMethodPingDatabase:
				go healthProtocol.handlePingDatabase(packet)
				break
			case HealthMethodRunSanityCheck:
				go healthProtocol.handleRunSanityCheck(packet)
				break
			case HealthMethodFixSanityErrors:
				go healthProtocol.handleFixSanityErrors(packet)
				break
			default:
				fmt.Printf("Unsupported Health method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (healthProtocol *HealthProtocol) handlePingDaemon(packet nex.PacketInterface) {
	if healthProtocol.PingDaemonHandler == nil {
		fmt.Println("[Warning] HealthProtocol::PingDaemon not implemented")
		go respondNotImplemented(packet, HealthProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go healthProtocol.PingDaemonHandler(nil, client, callID)
}

func (healthProtocol *HealthProtocol) handlePingDatabase(packet nex.PacketInterface) {
	if healthProtocol.PingDatabaseHandler == nil {
		fmt.Println("[Warning] HealthProtocol::PingDatabase not implemented")
		go respondNotImplemented(packet, HealthProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go healthProtocol.PingDatabaseHandler(nil, client, callID)
}

func (healthProtocol *HealthProtocol) handleRunSanityCheck(packet nex.PacketInterface) {
	if healthProtocol.RunSanityCheckHandler == nil {
		fmt.Println("[Warning] HealthProtocol::RunSanityCheck not implemented")
		go respondNotImplemented(packet, HealthProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go healthProtocol.RunSanityCheckHandler(nil, client, callID)
}

func (healthProtocol *HealthProtocol) handleFixSanityErrors(packet nex.PacketInterface) {
	if healthProtocol.FixSanityErrorsHandler == nil {
		fmt.Println("[Warning] HealthProtocol::FixSanityErrors not implemented")
		go respondNotImplemented(packet, HealthProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go healthProtocol.FixSanityErrorsHandler(nil, client, callID)
}

// PingDaemon sets the PingDaemon handler function
func (healthProtocol *HealthProtocol) PingDaemon(handler func(err error, client *nex.Client, callID uint32)) {
	healthProtocol.PingDaemonHandler = handler
}

// PingDatabase sets the PingDatabase handler function
func (healthProtocol *HealthProtocol) PingDatabase(handler func(err error, client *nex.Client, callID uint32)) {
	healthProtocol.PingDatabaseHandler = handler
}

// RunSanityCheck sets the RunSanityCheck handler function
func (healthProtocol *HealthProtocol) RunSanityCheck(handler func(err error, client *nex.Client, callID uint32)) {
	healthProtocol.RunSanityCheckHandler = handler
}

// FixSanityErrors sets the FixSanityErrors handler function
func (healthProtocol *HealthProtocol) FixSanityErrors(handler func(err error, client *nex.Client, callID uint32)) {
	healthProtocol.FixSanityErrorsHandler = handler
}

// NewHealthProtocol returns a new HealthProtocol
func NewHealthProtocol(server *nex.Server) *HealthProtocol {
	healthProtocol := &HealthProtocol{server: server}

	healthProtocol.Setup()

	return healthProtocol
}
