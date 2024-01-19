// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetcompetitionRankingScoreByPeriodList(packet nex.PacketInterface) {
	if protocol.GetcompetitionRankingScoreByPeriodList == nil {
		globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.GetcompetitionRankingScoreByPeriodList(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
