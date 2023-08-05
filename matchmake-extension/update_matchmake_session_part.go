// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// UpdateMatchmakeSessionPart sets the UpdateMatchmakeSessionPart handler function
func (protocol *Protocol) UpdateMatchmakeSessionPart(handler func(err error, client *nex.Client, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) uint32) {
	protocol.updateMatchmakeSessionPartHandler = handler
}

func (protocol *Protocol) handleUpdateMatchmakeSessionPart(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateMatchmakeSessionPartHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSessionPart not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	updateMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making_types.NewUpdateMatchmakeSessionParam())
	if err != nil {
		errorCode = protocol.updateMatchmakeSessionPartHandler(fmt.Errorf("Failed to read updateMatchmakeSessionParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateMatchmakeSessionPartHandler(nil, client, callID, updateMatchmakeSessionParam.(*match_making_types.UpdateMatchmakeSessionParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
