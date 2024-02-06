// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestProbeInitiation(packet nex.PacketInterface) {
	var err error

	if protocol.RequestProbeInitiation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::RequestProbeInitiation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	urlTargetList := types.NewList[*types.StationURL]()
	urlTargetList.Type = types.NewStationURL("")
	err = urlTargetList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestProbeInitiation(fmt.Errorf("Failed to read urlTargetList from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestProbeInitiation(nil, packet, callID, urlTargetList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
