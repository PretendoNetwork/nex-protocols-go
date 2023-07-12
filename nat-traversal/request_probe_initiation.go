// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiation sets the RequestProbeInitiation handler function
func (protocol *NATTraversalProtocol) RequestProbeInitiation(handler func(err error, client *nex.Client, callID uint32, urlTargetList []*nex.StationURL)) {
	protocol.RequestProbeInitiationHandler = handler
}

func (protocol *NATTraversalProtocol) handleRequestProbeInitiation(packet nex.PacketInterface) {
	if protocol.RequestProbeInitiationHandler == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlTargetList, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.RequestProbeInitiationHandler(fmt.Errorf("Failed to read urlTargetList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.RequestProbeInitiationHandler(nil, client, callID, urlTargetList)
}
