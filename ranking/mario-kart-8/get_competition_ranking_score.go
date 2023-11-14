// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCompetitionRankingScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetCompetitionRankingScore == nil {
		globals.Logger.Warning("RankingMarioKart8::GetCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingMarioKart8::GetCompetitionRankingScore STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.GetCompetitionRankingScore(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
