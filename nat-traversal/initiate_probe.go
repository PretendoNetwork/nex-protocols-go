// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InitiateProbe sets the InitiateProbe handler function
func (protocol *NATTraversalProtocol) InitiateProbe(handler func(err error, client *nex.Client, callID uint32, urlStationToProbe *nex.StationURL)) {
	protocol.InitiateProbeHandler = handler
}

func (protocol *NATTraversalProtocol) handleInitiateProbe(packet nex.PacketInterface) {
	if protocol.InitiateProbeHandler == nil {
		globals.Logger.Warning("NATTraversal::InitiateProbe not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlStationToProbe, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.InitiateProbeHandler(fmt.Errorf("Failed to read urlStationToProbe from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.InitiateProbeHandler(nil, client, callID, urlStationToProbe)
}
