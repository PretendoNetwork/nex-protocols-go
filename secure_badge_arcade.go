package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// SecureBadgeArcadeProtocolID is the protocol ID for the Secure Connection (Badge Arcade) protocol. ID is the same as the Secure Connection Protocol
	SecureBadgeArcadeProtocolID = 0xB

	// SecureBadgeArcadeMethodGetMaintenanceStatus is the method ID for GetMaintenanceStatus
	SecureBadgeArcadeMethodGetMaintenanceStatus = 0x9
)

// SecureBadgeArcadeProtocol handles the Secure Connection (Badge Arcade) nex protocol. Embeds SecureProtocol
type SecureBadgeArcadeProtocol struct {
	server *nex.Server
	SecureProtocol
	GetMaintenanceStatusHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (secureBadgeArcadeProtocol *SecureBadgeArcadeProtocol) Setup() {
	nexServer := secureBadgeArcadeProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if SecureBadgeArcadeProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case SecureMethodRegister:
				go secureBadgeArcadeProtocol.handleRegister(packet)
			case SecureBadgeArcadeMethodGetMaintenanceStatus:
				go secureBadgeArcadeProtocol.handleGetMaintenanceStatus(packet)
			default:
				go respondNotImplemented(packet, SecureBadgeArcadeProtocolID)
				fmt.Printf("Unsupported SecureBadgeArcade method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetMaintenanceStatus sets the GetMaintenanceStatus function
func (secureBadgeArcadeProtocol *SecureBadgeArcadeProtocol) GetMaintenanceStatus(handler func(err error, client *nex.Client, callID uint32)) {
	secureBadgeArcadeProtocol.GetMaintenanceStatusHandler = handler
}

func (secureBadgeArcadeProtocol *SecureBadgeArcadeProtocol) handleGetMaintenanceStatus(packet nex.PacketInterface) {
	if secureBadgeArcadeProtocol.GetMaintenanceStatusHandler == nil {
		logger.Warning("SecureBadgeArcadeProtocol::GetMaintenanceStatus not implemented")
		go respondNotImplemented(packet, SecureBadgeArcadeProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go secureBadgeArcadeProtocol.GetMaintenanceStatusHandler(nil, client, callID)
}

// NewSecureBadgeArcadeProtocol returns a new SecureBadgeArcadeProtocol
func NewSecureBadgeArcadeProtocol(server *nex.Server) *SecureBadgeArcadeProtocol {
	secureBadgeArcadeProtocol := &SecureBadgeArcadeProtocol{server: server}
	secureBadgeArcadeProtocol.SecureProtocol.server = server

	secureBadgeArcadeProtocol.Setup()

	return secureBadgeArcadeProtocol
}
