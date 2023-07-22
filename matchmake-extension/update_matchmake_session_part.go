// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// UpdateMatchmakeSessionPart sets the UpdateMatchmakeSessionPart handler function
func (protocol *MatchmakeExtensionProtocol) UpdateMatchmakeSessionPart(handler func(err error, client *nex.Client, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam)) {
	protocol.updateMatchmakeSessionPartHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleUpdateMatchmakeSessionPart(packet nex.PacketInterface) {
	if protocol.updateMatchmakeSessionPartHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSessionPart not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	updateMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making_types.NewUpdateMatchmakeSessionParam())
	if err != nil {
		go protocol.updateMatchmakeSessionPartHandler(fmt.Errorf("Failed to read updateMatchmakeSessionParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateMatchmakeSessionPartHandler(nil, client, callID, updateMatchmakeSessionParam.(*match_making_types.UpdateMatchmakeSessionParam))
}
