// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// StartRound sets the StartRound handler function
func (protocol *MatchmakeRefereeProtocol) StartRound(handler func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam)) {
	protocol.StartRoundHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleStartRound(packet nex.PacketInterface) {
	if protocol.StartRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::StartRound not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStartRoundParam())
	if err != nil {
		go protocol.StartRoundHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.StartRoundHandler(nil, client, callID, param.(*matchmake_referee_types.MatchmakeRefereeStartRoundParam))
}
