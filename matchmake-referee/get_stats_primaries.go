// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsPrimaries sets the GetStatsPrimaries handler function
func (protocol *Protocol) GetStatsPrimaries(handler func(err error, packet nex.PacketInterface, callID uint32, targets []*matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32) {
	protocol.getStatsPrimariesHandler = handler
}

func (protocol *Protocol) handleGetStatsPrimaries(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStatsPrimariesHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimaries not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		errorCode = protocol.getStatsPrimariesHandler(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStatsPrimariesHandler(nil, packet, callID, targets.([]*matchmake_referee_types.MatchmakeRefereeStatsTarget))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
