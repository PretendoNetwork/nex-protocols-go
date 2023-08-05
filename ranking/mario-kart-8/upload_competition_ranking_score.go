// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
)

// UploadCompetitionRankingScore sets the UploadCompetitionRankingScore handler function
func (protocol *Protocol) UploadCompetitionRankingScore(handler func(err error, client *nex.Client, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam) uint32) {
	protocol.uploadCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.uploadCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingMarioKart8::UploadCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(ranking_mario_kart8_types.NewCompetitionRankingUploadScoreParam())
	if err != nil {
		errorCode = protocol.uploadCompetitionRankingScoreHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.uploadCompetitionRankingScoreHandler(nil, client, callID, param.(*ranking_mario_kart8_types.CompetitionRankingUploadScoreParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
