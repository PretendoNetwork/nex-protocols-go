// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionNoHolderNoResultRange(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.BrowseMatchmakeSessionNoHolderNoResultRange == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionNoHolderNoResultRange not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	searchCriteria := match_making_types.NewMatchmakeSessionSearchCriteria()
	err = searchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.BrowseMatchmakeSessionNoHolderNoResultRange(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.BrowseMatchmakeSessionNoHolderNoResultRange(nil, packet, callID, searchCriteria)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
