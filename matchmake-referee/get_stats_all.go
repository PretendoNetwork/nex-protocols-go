// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsAll sets the GetStatsAll handler function
func (protocol *Protocol) GetStatsAll(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32) {
	protocol.getStatsAllHandler = handler
}

func (protocol *Protocol) handleGetStatsAll(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStatsAllHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsAll not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		errorCode = protocol.getStatsAllHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStatsAllHandler(nil, packet, callID, target.(*matchmake_referee_types.MatchmakeRefereeStatsTarget))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
