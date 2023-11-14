// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportNATProperties(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ReportNATProperties == nil {
		globals.Logger.Warning("NATTraversal::ReportNATProperties not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	natmapping, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read natmapping from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	natfiltering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read natfiltering from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rtt, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.ReportNATProperties(nil, packet, callID, natmapping, natfiltering, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
