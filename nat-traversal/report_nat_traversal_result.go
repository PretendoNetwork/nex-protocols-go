// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResult sets the ReportNATTraversalResult handler function
func (protocol *Protocol) ReportNATTraversalResult(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32) uint32) {
	protocol.reportNATTraversalResultHandler = handler
}

func (protocol *Protocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportNATTraversalResultHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResult not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	natTraversalVersion := protocol.Server.NATTraversalProtocolVersion()

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), client, callID, 0, false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), client, callID, 0, false, 0)
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
			errorCode = protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, false, 0)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}
	}

	errorCode = protocol.reportNATTraversalResultHandler(nil, client, callID, cid, result, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
