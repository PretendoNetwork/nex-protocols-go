// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ReportNATTraversalResult == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResult not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	natTraversalVersion := protocol.Server.NATTraversalProtocolVersion()

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, 0, false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, 0, false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	var rtt uint32 = 0

	// TODO - Is this the right version?
	if natTraversalVersion.GreaterOrEqual("3.0.0") {
		rtt, err = parametersStream.ReadUInt32LE()
		if err != nil {
			_, errorCode = protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, 0, false, 0)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}
	}

	rmcMessage, errorCode := protocol.ReportNATTraversalResult(nil, packet, callID, cid, result, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
