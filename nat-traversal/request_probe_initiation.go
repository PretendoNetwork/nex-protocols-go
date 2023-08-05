// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiation sets the RequestProbeInitiation handler function
func (protocol *Protocol) RequestProbeInitiation(handler func(err error, client *nex.Client, callID uint32, urlTargetList []*nex.StationURL) uint32) {
	protocol.requestProbeInitiationHandler = handler
}

func (protocol *Protocol) handleRequestProbeInitiation(packet nex.PacketInterface) {
	if protocol.requestProbeInitiationHandler == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlTargetList, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.requestProbeInitiationHandler(fmt.Errorf("Failed to read urlTargetList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.requestProbeInitiationHandler(nil, client, callID, urlTargetList)
}
