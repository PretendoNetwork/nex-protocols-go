// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetOrCreateStats sets the GetOrCreateStats handler function
func (protocol *Protocol) GetOrCreateStats(handler func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) uint32) {
	protocol.getOrCreateStatsHandler = handler
}

func (protocol *Protocol) handleGetOrCreateStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getOrCreateStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetOrCreateStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsInitParam())
	if err != nil {
		errorCode = protocol.getOrCreateStatsHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getOrCreateStatsHandler(nil, client, callID, param.(*matchmake_referee_types.MatchmakeRefereeStatsInitParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
