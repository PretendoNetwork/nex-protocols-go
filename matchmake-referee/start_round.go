// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

// StartRound sets the StartRound handler function
func (protocol *Protocol) StartRound(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) uint32) {
	protocol.startRoundHandler = handler
}

func (protocol *Protocol) handleStartRound(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.startRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::StartRound not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStartRoundParam())
	if err != nil {
		errorCode = protocol.startRoundHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.startRoundHandler(nil, packet, callID, param.(*matchmake_referee_types.MatchmakeRefereeStartRoundParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
