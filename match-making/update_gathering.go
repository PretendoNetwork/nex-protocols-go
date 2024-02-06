// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateGathering(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateGathering == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::UpdateGathering not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	anyGathering := types.NewAnyDataHolder()
	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateGathering(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateGathering(nil, packet, callID, anyGathering)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
