// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var err error

	if protocol.GetStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	category := types.NewPrimitiveU32(0)
	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	orderParam := ranking_types.NewRankingOrderParam()
	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	flags := types.NewPrimitiveU32(0)
	err = flags.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read flags from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStats(nil, packet, callID, category, orderParam, flags)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
