// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (protocol *Protocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)) {
	protocol.requestProbeInitiationExtHandler = handler
}

func (protocol *Protocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if protocol.reportNATPropertiesHandler == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiationExt not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targetList, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.requestProbeInitiationExtHandler(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	stationToProbe, err := parametersStream.ReadString()
	if err != nil {
		go protocol.requestProbeInitiationExtHandler(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	go protocol.requestProbeInitiationExtHandler(nil, client, callID, targetList, stationToProbe)
}
