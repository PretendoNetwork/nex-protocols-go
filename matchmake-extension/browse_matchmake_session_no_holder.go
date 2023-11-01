// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// BrowseMatchmakeSessionNoHolder sets the BrowseMatchmakeSessionNoHolder handler function
func (protocol *Protocol) BrowseMatchmakeSessionNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32) {
	protocol.browseMatchmakeSessionNoHolderHandler = handler
}

func (protocol *Protocol) handleBrowseMatchmakeSessionNoHolder(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.browseMatchmakeSessionNoHolderHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::BrowseMatchmakeSessionNoHolder not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	searchCriteria, err := parametersStream.ReadStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		errorCode = protocol.browseMatchmakeSessionNoHolderHandler(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.browseMatchmakeSessionNoHolderHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.browseMatchmakeSessionNoHolderHandler(nil, packet, callID, searchCriteria.(*match_making_types.MatchmakeSessionSearchCriteria), resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
