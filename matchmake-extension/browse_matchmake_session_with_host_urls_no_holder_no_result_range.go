// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	searchCriteria, err := parametersStream.ReadStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		_, errorCode = protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(nil, packet, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
