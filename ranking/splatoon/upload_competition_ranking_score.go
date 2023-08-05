// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCompetitionRankingScore sets the UploadCompetitionRankingScore handler function
func (protocol *Protocol) UploadCompetitionRankingScore(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.uploadCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.uploadCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingSplatoon::UploadCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::UploadCompetitionRankingScore STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.uploadCompetitionRankingScoreHandler(nil, client, callID, packet.Payload())
}
