// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLs(packet nex.PacketInterface) {
	if protocol.BrowseMatchmakeSessionWithHostURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::BrowseMatchmakeSessionWithHostURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	searchCriteria := match_making_types.NewMatchmakeSessionSearchCriteria()
	resultRange := types.NewResultRange()

	var err error

	err = searchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(nil, packet, callID, searchCriteria, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
