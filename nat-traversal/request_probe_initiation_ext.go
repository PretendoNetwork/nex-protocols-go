// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (protocol *Protocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string) uint32) {
	protocol.requestProbeInitiationExtHandler = handler
}

func (protocol *Protocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportNATPropertiesHandler == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiationExt not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targetList, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.requestProbeInitiationExtHandler(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), client, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	stationToProbe, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.requestProbeInitiationExtHandler(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), client, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestProbeInitiationExtHandler(nil, client, callID, targetList, stationToProbe)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
