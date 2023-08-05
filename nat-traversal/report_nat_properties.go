// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATProperties sets the ReportNATProperties handler function
func (protocol *Protocol) ReportNATProperties(handler func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32) uint32) {
	protocol.reportNATPropertiesHandler = handler
}

func (protocol *Protocol) handleReportNATProperties(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportNATPropertiesHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATProperties not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	natmapping, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportNATPropertiesHandler(fmt.Errorf("Failed to read natmapping from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	natfiltering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportNATPropertiesHandler(fmt.Errorf("Failed to read natfiltering from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rtt, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportNATPropertiesHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.reportNATPropertiesHandler(nil, client, callID, natmapping, natfiltering, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
