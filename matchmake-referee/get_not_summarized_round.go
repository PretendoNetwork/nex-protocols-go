// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotSummarizedRound sets the GetNotSummarizedRound handler function
func (protocol *MatchmakeRefereeProtocol) GetNotSummarizedRound(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getNotSummarizedRoundHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetNotSummarizedRound(packet nex.PacketInterface) {
	if protocol.getNotSummarizedRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetNotSummarizedRound not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getNotSummarizedRoundHandler(nil, client, callID)
}
