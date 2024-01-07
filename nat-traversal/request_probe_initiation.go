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
	var errorCode uint32

	if protocol.RequestProbeInitiation == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiation not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
		_, errorCode = protocol.RequestProbeInitiation(fmt.Errorf("Failed to read urlTargetList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestProbeInitiation(nil, packet, callID, urlTargetList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
