// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
)

// UploadCompetitionRankingScore sets the UploadCompetitionRankingScore handler function
func (protocol *Protocol) UploadCompetitionRankingScore(handler func(err error, packet nex.PacketInterface, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam) uint32) {
	protocol.uploadCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.uploadCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingMarioKart8::UploadCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(ranking_mario_kart8_types.NewCompetitionRankingUploadScoreParam())
	if err != nil {
		errorCode = protocol.uploadCompetitionRankingScoreHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.uploadCompetitionRankingScoreHandler(nil, packet, callID, param.(*ranking_mario_kart8_types.CompetitionRankingUploadScoreParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
