// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"encoding/hex"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// AutoMatchmakeWithSearchCriteriaPostpone sets the AutoMatchmakeWithSearchCriteriaPostpone handler function
func (protocol *Protocol) AutoMatchmakeWithSearchCriteriaPostpone(handler func(err error, client *nex.Client, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *nex.DataHolder, strMessage string) uint32) {
	protocol.autoMatchmakeWithSearchCriteriaPostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakeWithSearchCriteriaPostpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithSearchCriteriaPostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithSearchCriteriaPostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	globals.Logger.Info(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstSearchCriteria, err := parametersStream.ReadListStructure(match_making_types.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		go protocol.autoMatchmakeWithSearchCriteriaPostponeHandler(fmt.Errorf("Failed to read lstSearchCriteria from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.autoMatchmakeWithSearchCriteriaPostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.autoMatchmakeWithSearchCriteriaPostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	go protocol.autoMatchmakeWithSearchCriteriaPostponeHandler(nil, client, callID, lstSearchCriteria.([]*match_making_types.MatchmakeSessionSearchCriteria), anyGathering, strMessage)
}
