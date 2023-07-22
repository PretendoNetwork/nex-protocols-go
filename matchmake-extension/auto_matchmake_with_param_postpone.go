// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// AutoMatchmakeWithParam_Postpone sets the AutoMatchmakeWithParam_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmakeWithParam_Postpone(handler func(err error, client *nex.Client, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam)) {
	protocol.autoMatchmakeWithParam_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithParam_Postpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithParam_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithParam_Postpone not implemented")
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
		go protocol.autoMatchmakeWithParam_PostponeHandler(fmt.Errorf("Failed to read autoMatchmakeParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.autoMatchmakeWithParam_PostponeHandler(nil, client, callID, autoMatchmakeParam.(*match_making_types.AutoMatchmakeParam))
}
