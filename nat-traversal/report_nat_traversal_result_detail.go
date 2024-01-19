// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportNATTraversalResultDetail(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ReportNATTraversalResultDetail == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResultDetail not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	// TODO - The NEX server should add a NATTraversalProtocolVersion method
	matchmakingVersion := protocol.server.MatchMakingProtocolVersion()

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	cid := types.NewPrimitiveU32(0)
	err = cid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	result := types.NewPrimitiveBool(false)
	err = result.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	detail := types.NewPrimitiveS32(0)
	err = detail.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read detail from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rtt := types.NewPrimitiveU32(0)

	// TODO - Is this the right version?
	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		rttU32, err := parametersStream.ReadPrimitiveUInt32LE()
		if err != nil {
			_, errorCode = protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}

		rtt = types.NewPrimitiveU32(rttU32)
	}

	rmcMessage, errorCode := protocol.ReportNATTraversalResultDetail(nil, packet, callID, cid, result, detail, rtt)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
