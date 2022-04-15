package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MonitoringProtocolID is the protocol ID for the Monitoring protocol
	MonitoringProtocolID = 0x13

	// MonitoringMethodPingDaemon is the method ID for the method PingDaemon
	MonitoringMethodPingDaemon = 0x1

	// MonitoringMethodGetClusterMembers is the method ID for the method GetClusterMembers
	MonitoringMethodGetClusterMembers = 0x2
)

// MonitoringProtocol handles the Monitoring protocol
type MonitoringProtocol struct {
	server                   *nex.Server
	PingDaemonHandler        func(err error, client *nex.Client, callID uint32)
	GetClusterMembersHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (monitoringProtocol *MonitoringProtocol) Setup() {
	nexServer := monitoringProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MonitoringProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MonitoringMethodPingDaemon:
				go monitoringProtocol.handlePingDaemon(packet)
				break
			case MonitoringMethodGetClusterMembers:
				go monitoringProtocol.handleGetClusterMembers(packet)
				break
			default:
				fmt.Printf("Unsupported Monitoring method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (monitoringProtocol *MonitoringProtocol) handlePingDaemon(packet nex.PacketInterface) {
	if monitoringProtocol.PingDaemonHandler == nil {
		logger.Warning("MonitoringProtocol::PingDaemon not implemented")
		go respondNotImplemented(packet, MonitoringProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go monitoringProtocol.PingDaemonHandler(nil, client, callID)
}

func (monitoringProtocol *MonitoringProtocol) handleGetClusterMembers(packet nex.PacketInterface) {
	if monitoringProtocol.GetClusterMembersHandler == nil {
		logger.Warning("MonitoringProtocol::GetClusterMembers not implemented")
		go respondNotImplemented(packet, MonitoringProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go monitoringProtocol.GetClusterMembersHandler(nil, client, callID)
}

// PingDaemon sets the PingDaemon handler function
func (monitoringProtocol *MonitoringProtocol) PingDaemon(handler func(err error, client *nex.Client, callID uint32)) {
	monitoringProtocol.PingDaemonHandler = handler
}

// GetClusterMembers sets the GetClusterMembers handler function
func (monitoringProtocol *MonitoringProtocol) GetClusterMembers(handler func(err error, client *nex.Client, callID uint32)) {
	monitoringProtocol.GetClusterMembersHandler = handler
}

// NewMonitoringProtocol returns a new MonitoringProtocol
func NewMonitoringProtocol(server *nex.Server) *MonitoringProtocol {
	monitoringProtocol := &MonitoringProtocol{server: server}

	monitoringProtocol.Setup()

	return monitoringProtocol
}
