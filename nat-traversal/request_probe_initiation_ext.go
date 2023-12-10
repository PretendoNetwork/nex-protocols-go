// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ReportNATProperties == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiationExt not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	targetList, err := parametersStream.ReadListString()
	if err != nil {
		_, errorCode = protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	stationToProbe, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestProbeInitiationExt(nil, packet, callID, targetList, stationToProbe)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
