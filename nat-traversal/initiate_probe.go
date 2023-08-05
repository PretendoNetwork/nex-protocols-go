// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InitiateProbe sets the InitiateProbe handler function
func (protocol *Protocol) InitiateProbe(handler func(err error, client *nex.Client, callID uint32, urlStationToProbe *nex.StationURL) uint32) {
	protocol.initiateProbeHandler = handler
}

func (protocol *Protocol) handleInitiateProbe(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.initiateProbeHandler == nil {
		globals.Logger.Warning("NATTraversal::InitiateProbe not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlStationToProbe, err := parametersStream.ReadStationURL()
	if err != nil {
		errorCode = protocol.initiateProbeHandler(fmt.Errorf("Failed to read urlStationToProbe from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.initiateProbeHandler(nil, client, callID, urlStationToProbe)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
