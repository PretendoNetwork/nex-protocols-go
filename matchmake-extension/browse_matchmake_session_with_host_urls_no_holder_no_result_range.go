// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange sets the BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange handler function
func (protocol *MatchmakeExtensionProtocol) BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria)) {
	protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet nex.PacketInterface) {
	if protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	searchCriteria, err := parametersStream.ReadStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		go protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.browseMatchmakeSessionWithHostURLsNoHolderNoResultRangeHandler(nil, client, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria))
}