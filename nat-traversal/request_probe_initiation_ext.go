// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if protocol.RequestProbeInitiationExt == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::RequestProbeInitiationExt not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var targetList types.List[types.String]
	var stationToProbe types.String

	var err error

	err = targetList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), packet, callID, targetList, stationToProbe)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = stationToProbe.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), packet, callID, targetList, stationToProbe)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestProbeInitiationExt(nil, packet, callID, targetList, stationToProbe)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
