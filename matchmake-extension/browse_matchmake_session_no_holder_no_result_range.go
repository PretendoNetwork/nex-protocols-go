// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionNoHolderNoResultRange sets the BrowseMatchmakeSessionNoHolderNoResultRange handler function
func (protocol *Protocol) BrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) uint32) {
	protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler = handler
}

func (protocol *Protocol) handleBrowseMatchmakeSessionNoHolderNoResultRange(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionNoHolderNoResultRange not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	searchCriteria, err := parametersStream.ReadStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		errorCode = protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler(nil, packet, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
