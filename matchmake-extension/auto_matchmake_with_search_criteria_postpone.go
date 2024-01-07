// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleAutoMatchmakeWithSearchCriteriaPostpone(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AutoMatchmakeWithSearchCriteriaPostpone == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithSearchCriteriaPostpone not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstSearchCriteria, err := nex.StreamReadListStructure(parametersStream, match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read lstSearchCriteria from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	anyGathering := types.NewAnyDataHolder()
	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage := types.NewString("")
	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AutoMatchmakeWithSearchCriteriaPostpone(nil, packet, callID, lstSearchCriteria, anyGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
