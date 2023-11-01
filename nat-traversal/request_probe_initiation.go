// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestProbeInitiation sets the RequestProbeInitiation handler function
func (protocol *Protocol) RequestProbeInitiation(handler func(err error, packet nex.PacketInterface, callID uint32, urlTargetList []*nex.StationURL) uint32) {
	protocol.requestProbeInitiationHandler = handler
}

func (protocol *Protocol) handleRequestProbeInitiation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestProbeInitiationHandler == nil {
		globals.Logger.Warning("NATTraversal::RequestProbeInitiation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlTargetList, err := parametersStream.ReadListStationURL()
	if err != nil {
		errorCode = protocol.requestProbeInitiationHandler(fmt.Errorf("Failed to read urlTargetList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestProbeInitiationHandler(nil, packet, callID, urlTargetList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
