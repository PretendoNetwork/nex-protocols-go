// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsPrimaries sets the GetStatsPrimaries handler function
func (protocol *MatchmakeRefereeProtocol) GetStatsPrimaries(handler func(err error, client *nex.Client, callID uint32, targets []*matchmake_referee_types.MatchmakeRefereeStatsTarget)) {
	protocol.getStatsPrimariesHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetStatsPrimaries(packet nex.PacketInterface) {
	if protocol.getStatsPrimariesHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimaries not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		go protocol.getStatsPrimariesHandler(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getStatsPrimariesHandler(nil, client, callID, targets.([]*matchmake_referee_types.MatchmakeRefereeStatsTarget))
}
