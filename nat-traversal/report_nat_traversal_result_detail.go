// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReportNATTraversalResultDetail(packet nex.PacketInterface) {
	if protocol.ReportNATTraversalResultDetail == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::ReportNATTraversalResultDetail not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()
	natTraversalVersion := endpoint.LibraryVersions().NATTraversal

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var cid types.UInt32
	var result types.Bool
	var detail types.Int32
	var rtt types.UInt32

	var err error

	err = cid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, cid, result, detail, rtt)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = result.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, cid, result, detail, rtt)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = detail.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read detail from parameters. %s", err.Error()), packet, callID, cid, result, detail, rtt)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	// TODO - Is this the right version?
	if natTraversalVersion.GreaterOrEqual("3.0.0") {
		err = rtt.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.ReportNATTraversalResultDetail(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, cid, result, detail, rtt)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	}

	rmcMessage, rmcError := protocol.ReportNATTraversalResultDetail(nil, packet, callID, cid, result, detail, rtt)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
