// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleInitiateProbe(packet nex.PacketInterface) {
	if protocol.InitiateProbe == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::InitiateProbe not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	urlStationToProbe := types.NewStationURL("")

	err := urlStationToProbe.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InitiateProbe(fmt.Errorf("Failed to read urlStationToProbe from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.InitiateProbe(nil, packet, callID, urlStationToProbe)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
