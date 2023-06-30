// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATProperties sets the ReportNATProperties handler function
func (protocol *NATTraversalProtocol) ReportNATProperties(handler func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)) {
	protocol.ReportNATPropertiesHandler = handler
}

func (protocol *NATTraversalProtocol) handleReportNATProperties(packet nex.PacketInterface) {
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

	natmapping, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.ReportNATPropertiesHandler(fmt.Errorf("Failed to read natmapping from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	natfiltering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.ReportNATPropertiesHandler(fmt.Errorf("Failed to read natfiltering from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	rtt, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.ReportNATPropertiesHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	go protocol.ReportNATPropertiesHandler(nil, client, callID, natmapping, natfiltering, rtt)
}
