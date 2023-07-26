// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// CreateStats sets the CreateStats handler function
func (protocol *MatchmakeRefereeProtocol) CreateStats(handler func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam)) {
	protocol.CreateStatsHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleCreateStats(packet nex.PacketInterface) {
	if protocol.CreateStatsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::CreateStats not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsInitParam())
	if err != nil {
		go protocol.CreateStatsHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CreateStatsHandler(nil, client, callID, param.(*matchmake_referee_types.MatchmakeRefereeStatsInitParam))
}
