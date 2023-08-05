// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange sets the BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange handler function
func (protocol *Protocol) BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) uint32) {
	protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler = handler
}

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	searchCriteria, err := parametersStream.ReadStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		errorCode = protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler(nil, client, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
