// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionWithHostURLs sets the BrowseMatchmakeSessionWithHostURLs handler function
func (protocol *Protocol) BrowseMatchmakeSessionWithHostURLs(handler func(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32) {
	protocol.browseMatchmakeSessionWithHostURLsHandler = handler
}

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLs(packet nex.PacketInterface) {
	if protocol.browseMatchmakeSessionWithHostURLsHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionWithHostURLs not implemented")
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
		go protocol.browseMatchmakeSessionWithHostURLsHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.browseMatchmakeSessionWithHostURLsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.browseMatchmakeSessionWithHostURLsHandler(nil, client, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria), resultRange.(*nex.ResultRange))
}
