// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRoundParticipants sets the GetRoundParticipants handler function
func (protocol *Protocol) GetRoundParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) uint32) {
	protocol.getRoundParticipantsHandler = handler
}

func (protocol *Protocol) handleGetRoundParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRoundParticipantsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetRoundParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getRoundParticipantsHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRoundParticipantsHandler(nil, packet, callID, roundID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
