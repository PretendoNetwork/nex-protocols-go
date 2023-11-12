// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResultDetail sets the ReportNATTraversalResultDetail handler function
func (protocol *Protocol) ReportNATTraversalResultDetail(handler func(err error, packet nex.PacketInterface, callID uint32, cid uint32, result bool, detail int32, rtt uint32) uint32) {
	protocol.reportNATTraversalResultDetailHandler = handler
}

func (protocol *Protocol) handleReportNATTraversalResultDetail(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportNATTraversalResultDetailHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResultDetail not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	// TODO - The NEX server should add a NATTraversalProtocolVersion method
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, 0, false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, 0, false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	detail, err := parametersStream.ReadInt32LE()
	if err != nil {
		errorCode = protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read detail from parameters. %s", err.Error()), packet, callID, 0, false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	var rtt uint32 = 0

	// TODO - Is this the right version?
	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		rtt, err = parametersStream.ReadUInt32LE()
		if err != nil {
			errorCode = protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, 0, false, 0, 0)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}
	}

	errorCode = protocol.reportNATTraversalResultDetailHandler(nil, packet, callID, cid, result, detail, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
