// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetcompetitionRankingScoreByPeriodList sets the GetcompetitionRankingScoreByPeriodList handler function
func (protocol *Protocol) GetcompetitionRankingScoreByPeriodList(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getcompetitionRankingScoreByPeriodListHandler = handler
}

func (protocol *Protocol) handleGetcompetitionRankingScoreByPeriodList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getcompetitionRankingScoreByPeriodListHandler == nil {
		globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("RankingSplatoon::GetcompetitionRankingScoreByPeriodList STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getcompetitionRankingScoreByPeriodListHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
