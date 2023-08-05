// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// AutoMatchmakeWithParamPostpone sets the AutoMatchmakeWithParamPostpone handler function
func (protocol *Protocol) AutoMatchmakeWithParamPostpone(handler func(err error, client *nex.Client, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) uint32) {
	protocol.autoMatchmakeWithParamPostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakeWithParamPostpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithParamPostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithParamPostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	autoMatchmakeParam, err := parametersStream.ReadStructure(match_making_types.NewAutoMatchmakeParam())
	if err != nil {
		go protocol.autoMatchmakeWithParamPostponeHandler(fmt.Errorf("Failed to read autoMatchmakeParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.autoMatchmakeWithParamPostponeHandler(nil, client, callID, autoMatchmakeParam.(*match_making_types.AutoMatchmakeParam))
}
