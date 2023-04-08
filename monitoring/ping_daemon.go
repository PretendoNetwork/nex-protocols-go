package monitoring

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDaemon sets the PingDaemon handler function
func (protocol *MonitoringProtocol) PingDaemon(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.PingDaemonHandler = handler
}

func (protocol *MonitoringProtocol) HandlePingDaemon(packet nex.PacketInterface) {
	if protocol.PingDaemonHandler == nil {
		globals.Logger.Warning("Monitoring::PingDaemon not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.PingDaemonHandler(nil, client, callID)
}
