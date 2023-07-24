// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRound sets the GetRound handler function
func (protocol *MatchmakeRefereeProtocol) GetRound(handler func(err error, client *nex.Client, callID uint32, roundId uint64)) {
	protocol.GetRoundHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetRound(packet nex.PacketInterface) {
	if protocol.GetRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetRound not implemented")
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
		go protocol.GetRoundHandler(fmt.Errorf("Failed to read roundId from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetRoundHandler(nil, client, callID, roundId)
}
