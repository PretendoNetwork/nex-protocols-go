// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// RequestMatchmaking sets the RequestMatchmaking handler function
func (protocol *Protocol) RequestMatchmaking(handler func(err error, client *nex.Client, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam)) {
	protocol.requestMatchmakingHandler = handler
}

func (protocol *Protocol) handleRequestMatchmaking(packet nex.PacketInterface) {
	if protocol.requestMatchmakingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::RequestMatchmaking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	autoMatchmakeParam, err := parametersStream.ReadStructure(match_making_types.NewAutoMatchmakeParam())
	if err != nil {
		go protocol.requestMatchmakingHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.requestMatchmakingHandler(nil, client, callID, autoMatchmakeParam.(*match_making_types.AutoMatchmakeParam))
}
