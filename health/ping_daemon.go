package health

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDaemon sets the PingDaemon handler function
func (protocol *HealthProtocol) PingDaemon(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.PingDaemonHandler = handler
}

func (protocol *HealthProtocol) HandlePingDaemon(packet nex.PacketInterface) {
	if protocol.PingDaemonHandler == nil {
		globals.Logger.Warning("Health::PingDaemon not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.PingDaemonHandler(nil, client, callID)
}
