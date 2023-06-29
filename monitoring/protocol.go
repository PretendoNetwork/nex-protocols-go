package monitoring

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

// MonitoringProtocol handles the Monitoring protocol
type MonitoringProtocol struct {
	Server                   *nex.Server
	PingDaemonHandler        func(err error, client *nex.Client, callID uint32)
	GetClusterMembersHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *MonitoringProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *MonitoringProtocol) HandlePacket(packet nex.PacketInterface) {
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

// NewMonitoringProtocol returns a new MonitoringProtocol
func NewMonitoringProtocol(server *nex.Server) *MonitoringProtocol {
	protocol := &MonitoringProtocol{Server: server}

	protocol.Setup()

	return protocol
}
