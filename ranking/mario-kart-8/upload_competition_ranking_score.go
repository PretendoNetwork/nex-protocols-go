// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCompetitionRankingScore sets the UploadCompetitionRankingScore handler function
func (protocol *Protocol) UploadCompetitionRankingScore(handler func(err error, client *nex.Client, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam)) {
	protocol.uploadCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.uploadCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingMarioKart8::UploadCompetitionRankingScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(ranking_mario_kart8_types.NewCompetitionRankingUploadScoreParam())
	if err != nil {
		go protocol.uploadCompetitionRankingScoreHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.uploadCompetitionRankingScoreHandler(nil, client, callID, param.(*ranking_mario_kart8_types.CompetitionRankingUploadScoreParam))
}