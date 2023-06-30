// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (protocol *NATTraversalProtocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)) {
	protocol.RequestProbeInitiationExtHandler = handler
}

func (protocol *NATTraversalProtocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if protocol.ReportNATPropertiesHandler == nil {
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
		go protocol.RequestProbeInitiationExtHandler(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	stationToProbe, err := parametersStream.ReadString()
	if err != nil {
		go protocol.RequestProbeInitiationExtHandler(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	go protocol.RequestProbeInitiationExtHandler(nil, client, callID, targetList, stationToProbe)
}
