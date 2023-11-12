// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// CreateStats sets the CreateStats handler function
func (protocol *Protocol) CreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) uint32) {
	protocol.createStatsHandler = handler
}

func (protocol *Protocol) handleCreateStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::CreateStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsInitParam())
	if err != nil {
		errorCode = protocol.createStatsHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createStatsHandler(nil, packet, callID, param.(*matchmake_referee_types.MatchmakeRefereeStatsInitParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
