// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_splatoon_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon/types"
)

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.UploadCompetitionRankingScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "RankingSplatoon::UploadCompetitionRankingScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("RankingSplatoon::UploadCompetitionRankingScore STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID
	endpoint := request.Endpoint
	parameters := request.Parameters

	param := ranking_splatoon_types.NewCompetitionRankingUploadScoreParam()

	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadCompetitionRankingScore(fmt.Errorf("failed to read param from parameters. %s", err.Error()), packet, callID, param)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadCompetitionRankingScore(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
