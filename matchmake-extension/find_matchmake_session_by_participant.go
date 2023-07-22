// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// FindMatchmakeSessionByParticipant sets the FindMatchmakeSessionByParticipant handler function
func (protocol *MatchmakeExtensionProtocol) FindMatchmakeSessionByParticipant(handler func(err error, client *nex.Client, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam)) {
	protocol.findMatchmakeSessionByParticipantHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleFindMatchmakeSessionByParticipant(packet nex.PacketInterface) {
	if protocol.findMatchmakeSessionByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByParticipant not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(match_making_types.NewFindMatchmakeSessionByParticipantParam())
	if err != nil {
		go protocol.findMatchmakeSessionByParticipantHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.findMatchmakeSessionByParticipantHandler(nil, client, callID, param.(*match_making_types.FindMatchmakeSessionByParticipantParam))
}