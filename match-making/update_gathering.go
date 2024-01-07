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
	var errorCode uint32

	if protocol.UpdateGathering == nil {
		globals.Logger.Warning("MatchMaking::UpdateGathering not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	anyGathering := types.NewAnyDataHolder()
	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateGathering(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateGathering(nil, packet, callID, anyGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
