// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionNoHolderNoResultRange sets the BrowseMatchmakeSessionNoHolderNoResultRange handler function
func (protocol *MatchmakeExtensionProtocol) BrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria)) {
	protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleBrowseMatchmakeSessionNoHolderNoResultRange(packet nex.PacketInterface) {
	if protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionNoHolderNoResultRange not implemented")
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
		go protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.browseMatchmakeSessionNoHolderNoResultRangeHandler(nil, client, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria))
}
