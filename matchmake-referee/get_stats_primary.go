// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsPrimary sets the GetStatsPrimary handler function
func (protocol *Protocol) GetStatsPrimary(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32) {
	protocol.getStatsPrimaryHandler = handler
}

func (protocol *Protocol) handleGetStatsPrimary(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStatsPrimaryHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimary not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		errorCode = protocol.getStatsPrimaryHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStatsPrimaryHandler(nil, packet, callID, target.(*matchmake_referee_types.MatchmakeRefereeStatsTarget))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
