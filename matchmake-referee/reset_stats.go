// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetStats sets the ResetStats handler function
func (protocol *Protocol) ResetStats(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.resetStatsHandler = handler
}

func (protocol *Protocol) handleResetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.resetStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::ResetStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.resetStatsHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
