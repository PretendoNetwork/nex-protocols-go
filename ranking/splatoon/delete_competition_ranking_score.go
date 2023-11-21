// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCompetitionRankingScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteCompetitionRankingScore == nil {
		globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.DeleteCompetitionRankingScore(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
