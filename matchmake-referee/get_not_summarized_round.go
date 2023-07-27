// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotSummarizedRound sets the GetNotSummarizedRound handler function
func (protocol *Protocol) GetNotSummarizedRound(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getNotSummarizedRoundHandler = handler
}

func (protocol *Protocol) handleGetNotSummarizedRound(packet nex.PacketInterface) {
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
