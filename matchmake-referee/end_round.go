// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// EndRound sets the EndRound handler function
func (protocol *Protocol) EndRound(handler func(err error, client *nex.Client, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam)) {
	protocol.endRoundHandler = handler
}

func (protocol *Protocol) handleEndRound(packet nex.PacketInterface) {
	if protocol.endRoundHandler == nil {
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
		go protocol.endRoundHandler(fmt.Errorf("Failed to read endRoundParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.endRoundHandler(nil, client, callID, endRoundParam.(*matchmake_referee_types.MatchmakeRefereeEndRoundParam))
}
