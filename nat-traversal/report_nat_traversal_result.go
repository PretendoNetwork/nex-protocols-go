package nat_traversal

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResult sets the ReportNATTraversalResult handler function
func (protocol *NATTraversalProtocol) ReportNATTraversalResult(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)) {
	protocol.ReportNATTraversalResultHandler = handler
}

func (protocol *NATTraversalProtocol) HandleReportNATTraversalResult(packet nex.PacketInterface) {
	if protocol.ReportNATTraversalResultHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResult not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid := parametersStream.ReadUInt32LE()
	result := parametersStream.ReadBool()
	rtt := parametersStream.ReadUInt32LE()

	go protocol.ReportNATTraversalResultHandler(nil, client, callID, cid, result, rtt)
}
