// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// GetStatsPrimary sets the GetStatsPrimary handler function
func (protocol *Protocol) GetStatsPrimary(handler func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32) {
	protocol.getStatsPrimaryHandler = handler
}

func (protocol *Protocol) handleGetStatsPrimary(packet nex.PacketInterface) {
	if protocol.getStatsPrimaryHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimary not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		go protocol.getStatsPrimaryHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getStatsPrimaryHandler(nil, client, callID, target.(*matchmake_referee_types.MatchmakeRefereeStatsTarget))
}
