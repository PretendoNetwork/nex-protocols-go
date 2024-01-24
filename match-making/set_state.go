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
	var errorCode uint32

	if protocol.SetState == nil {
		globals.Logger.Warning("MatchMaking::SetState not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetState(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiNewState := types.NewPrimitiveU32(0)
	err = uiNewState.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetState(fmt.Errorf("Failed to read uiNewState from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SetState(nil, packet, callID, idGathering, uiNewState)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
