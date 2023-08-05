// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCompetitionRankingScore sets the GetCompetitionRankingScore handler function
func (protocol *Protocol) GetCompetitionRankingScore(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleGetCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.getCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingSplatoon::GetCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::GetCompetitionRankingScore STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getCompetitionRankingScoreHandler(nil, client, callID, packet.Payload())
}
