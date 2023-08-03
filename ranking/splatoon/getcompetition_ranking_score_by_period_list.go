// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetcompetitionRankingScoreByPeriodList sets the GetcompetitionRankingScoreByPeriodList handler function
func (protocol *Protocol) GetcompetitionRankingScoreByPeriodList(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getcompetitionRankingScoreByPeriodListHandler = handler
}

func (protocol *Protocol) handleGetcompetitionRankingScoreByPeriodList(packet nex.PacketInterface) {
	if protocol.getcompetitionRankingScoreByPeriodListHandler == nil {
		globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getcompetitionRankingScoreByPeriodListHandler(nil, client, callID, packet.Payload())
}