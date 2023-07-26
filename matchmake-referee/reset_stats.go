// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetStats sets the ResetStats handler function
func (protocol *MatchmakeRefereeProtocol) ResetStats(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.resetStatsHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleResetStats(packet nex.PacketInterface) {
	if protocol.resetStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::ResetStats not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.resetStatsHandler(nil, client, callID)
}
