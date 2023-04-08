package secure_connection_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection (Badge Arcade) protocol. ID is the same as the Secure Connection Protocol
	ProtocolID = 0xB

	// MethodGetMaintenanceStatus is the method ID for GetMaintenanceStatus
	MethodGetMaintenanceStatus = 0x9
)

var patchedMethods = []uint32{
	MethodGetMaintenanceStatus,
}

// SecureConnectionNintendoBadgeArcadeProtocol handles the Secure Connection (Nintendo Badge Arcade) nex protocol. Embeds SecureProtocol
type SecureConnectionNintendoBadgeArcadeProtocol struct {
	Server *nex.Server
	secure_connection.SecureConnectionProtocol
	GetMaintenanceStatusHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *SecureConnectionNintendoBadgeArcadeProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.SecureConnectionProtocol.HandlePacket(packet)
			}
		}
	})
}

func (protocol *SecureConnectionNintendoBadgeArcadeProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetMaintenanceStatus:
		go protocol.HandleGetMaintenanceStatus(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported SecureBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewSecureConnectionNintendoBadgeArcadeProtocol returns a new protocol
func NewSecureConnectionNintendoBadgeArcadeProtocol(server *nex.Server) *SecureConnectionNintendoBadgeArcadeProtocol {
	protocol := &SecureConnectionNintendoBadgeArcadeProtocol{Server: server}
	protocol.SecureConnectionProtocol.Server = server

	protocol.Setup()

	return protocol
}
