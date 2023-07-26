// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsAll sets the GetStatsAll handler function
func (protocol *MatchmakeRefereeProtocol) GetStatsAll(handler func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget)) {
	protocol.GetStatsAllHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetStatsAll(packet nex.PacketInterface) {
	if protocol.GetStatsAllHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsAll not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		go protocol.GetStatsAllHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetStatsAllHandler(nil, client, callID, target.(*matchmake_referee_types.MatchmakeRefereeStatsTarget))
}
