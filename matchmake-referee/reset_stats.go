// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetStats sets the ResetStats handler function
func (protocol *Protocol) ResetStats(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.resetStatsHandler = handler
}

func (protocol *Protocol) handleResetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.resetStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::ResetStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.resetStatsHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
