// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRoundParticipants sets the GetRoundParticipants handler function
func (protocol *MatchmakeRefereeProtocol) GetRoundParticipants(handler func(err error, client *nex.Client, callID uint32, roundID uint64)) {
	protocol.getRoundParticipantsHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetRoundParticipants(packet nex.PacketInterface) {
	if protocol.getRoundParticipantsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetRoundParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getRoundParticipantsHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getRoundParticipantsHandler(nil, client, callID, roundID)
}
