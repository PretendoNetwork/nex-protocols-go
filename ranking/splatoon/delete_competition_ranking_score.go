// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCompetitionRankingScore sets the DeleteCompetitionRankingScore handler function
func (protocol *Protocol) DeleteCompetitionRankingScore(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.deleteCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleDeleteCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.deleteCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.deleteCompetitionRankingScoreHandler(nil, client, callID, packet.Payload())
}
