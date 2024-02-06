// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleChangeAttributes(packet nex.PacketInterface) {
	var err error

	if protocol.ChangeAttributes == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::ChangeAttributes not implemented")

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
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	changeParam := ranking_types.NewRankingChangeAttributesParam()
	err = changeParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	uniqueID := types.NewPrimitiveU64(0)
	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ChangeAttributes(nil, packet, callID, category, changeParam, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
