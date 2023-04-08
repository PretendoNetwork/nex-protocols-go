package nat_traversal

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATProperties sets the ReportNATProperties handler function
func (protocol *NATTraversalProtocol) ReportNATProperties(handler func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)) {
	protocol.ReportNATPropertiesHandler = handler
}

func (protocol *NATTraversalProtocol) HandleReportNATProperties(packet nex.PacketInterface) {
	if protocol.ReportNATPropertiesHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATProperties not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	natmapping := parametersStream.ReadUInt32LE()
	natfiltering := parametersStream.ReadUInt32LE()
	rtt := parametersStream.ReadUInt32LE()

	go protocol.ReportNATPropertiesHandler(nil, client, callID, natmapping, natfiltering, rtt)
}
