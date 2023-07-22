// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSession sets the BrowseMatchmakeSession handler function
func (protocol *MatchmakeExtensionProtocol) BrowseMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange)) {
	protocol.browseMatchmakeSessionHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleBrowseMatchmakeSession(packet nex.PacketInterface) {
	if protocol.browseMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSession not implemented")
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
		go protocol.browseMatchmakeSessionHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.browseMatchmakeSessionHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.browseMatchmakeSessionHandler(nil, client, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria), resultRange.(*nex.ResultRange))
}
