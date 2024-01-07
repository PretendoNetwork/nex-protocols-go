// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportNATProperties(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ReportNATProperties == nil {
		globals.Logger.Warning("NATTraversal::ReportNATProperties not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	natmapping := types.NewPrimitiveU32(0)
	err = natmapping.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read natmapping from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	natfiltering := types.NewPrimitiveU32(0)
	err = natfiltering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read natfiltering from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rtt := types.NewPrimitiveU32(0)
	err = rtt.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATProperties(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ReportNATProperties(nil, packet, callID, natmapping, natfiltering, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
