// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStartRoundParam sets the GetStartRoundParam handler function
func (protocol *MatchmakeRefereeProtocol) GetStartRoundParam(handler func(err error, client *nex.Client, callID uint32, roundId uint64)) {
	protocol.GetStartRoundParamHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetStartRoundParam(packet nex.PacketInterface) {
	if protocol.GetStartRoundParamHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStartRoundParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundId, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.GetStartRoundParamHandler(fmt.Errorf("Failed to read roundId from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetStartRoundParamHandler(nil, client, callID, roundId)
}
