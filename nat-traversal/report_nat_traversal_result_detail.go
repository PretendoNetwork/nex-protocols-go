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

	if protocol.ReportNATTraversalResultDetail == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::ReportNATTraversalResultDetail not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

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
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	result := types.NewPrimitiveBool(false)
	err = result.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	detail := types.NewPrimitiveS32(0)
	err = detail.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read detail from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rtt := types.NewPrimitiveU32(0)

	// TODO - Is this the right version?
	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		rttU32, err := parametersStream.ReadPrimitiveUInt32LE()
		if err != nil {
			_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}

		rtt = types.NewPrimitiveU32(rttU32)
	}

	rmcMessage, rmcError := protocol.ReportNATTraversalResultDetail(nil, packet, callID, cid, result, detail, rtt)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
