// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetState(packet nex.PacketInterface) {
	var err error

	if protocol.SetState == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::SetState not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetState(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	uiNewState := types.NewPrimitiveU32(0)
	err = uiNewState.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetState(fmt.Errorf("Failed to read uiNewState from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetState(nil, packet, callID, idGathering, uiNewState)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
