// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
)

func (protocol *Protocol) handleGetCompetitionInfo(packet nex.PacketInterface) {
	if protocol.GetCompetitionInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "RankingMarioKart8::GetCompetitionInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	param := ranking_mario_kart8_types.NewCompetitionRankingInfoGetParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCompetitionInfo(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCompetitionInfo(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
