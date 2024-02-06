// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetRankingByPrincipalID(packet nex.PacketInterface) {
	var err error

	if protocol.GetRankingByPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetRankingByPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getParam := ranking2_types.NewRanking2GetParam()
	err = getParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	principalIDList := types.NewList[*types.PID]()
	principalIDList.Type = types.NewPID(0)
	err = principalIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingByPrincipalID(nil, packet, callID, getParam, principalIDList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
