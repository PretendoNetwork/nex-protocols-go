// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCompetitionRankingScore sets the DeleteCompetitionRankingScore handler function
func (protocol *Protocol) DeleteCompetitionRankingScore(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.deleteCompetitionRankingScoreHandler = handler
}

func (protocol *Protocol) handleDeleteCompetitionRankingScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteCompetitionRankingScoreHandler == nil {
		globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::DeleteCompetitionRankingScore STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.deleteCompetitionRankingScoreHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
