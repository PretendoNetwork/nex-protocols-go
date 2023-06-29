package monitoring

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetClusterMembers sets the GetClusterMembers handler function
func (protocol *MonitoringProtocol) GetClusterMembers(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetClusterMembersHandler = handler
}

func (protocol *MonitoringProtocol) handleGetClusterMembers(packet nex.PacketInterface) {
	if protocol.GetClusterMembersHandler == nil {
		globals.Logger.Warning("Monitoring::GetClusterMembers not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetClusterMembersHandler(nil, client, callID)
}
