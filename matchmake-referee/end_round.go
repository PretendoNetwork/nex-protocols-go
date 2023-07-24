// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// EndRound sets the EndRound handler function
func (protocol *MatchmakeRefereeProtocol) EndRound(handler func(err error, client *nex.Client, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam)) {
	protocol.EndRoundHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleEndRound(packet nex.PacketInterface) {
	if protocol.EndRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRound not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	endRoundParam, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeEndRoundParam())
	if err != nil {
		go protocol.EndRoundHandler(fmt.Errorf("Failed to read endRoundParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.EndRoundHandler(nil, client, callID, endRoundParam.(*matchmake_referee_types.MatchmakeRefereeEndRoundParam))
}
