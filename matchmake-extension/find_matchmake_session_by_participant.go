// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// FindMatchmakeSessionByParticipant sets the FindMatchmakeSessionByParticipant handler function
func (protocol *Protocol) FindMatchmakeSessionByParticipant(handler func(err error, client *nex.Client, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) uint32) {
	protocol.findMatchmakeSessionByParticipantHandler = handler
}

func (protocol *Protocol) handleFindMatchmakeSessionByParticipant(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findMatchmakeSessionByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByParticipant not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(match_making_types.NewFindMatchmakeSessionByParticipantParam())
	if err != nil {
		errorCode = protocol.findMatchmakeSessionByParticipantHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findMatchmakeSessionByParticipantHandler(nil, client, callID, param.(*match_making_types.FindMatchmakeSessionByParticipantParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
