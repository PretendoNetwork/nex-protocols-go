// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"encoding/hex"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// AutoMatchmakeWithSearchCriteria_Postpone sets the AutoMatchmakeWithSearchCriteria_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmakeWithSearchCriteria_Postpone(handler func(err error, client *nex.Client, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *nex.DataHolder, strMessage string)) {
	protocol.autoMatchmakeWithSearchCriteria_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithSearchCriteria_Postpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithSearchCriteria_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithSearchCriteria_Postpone not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
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
		go protocol.autoMatchmakeWithSearchCriteria_PostponeHandler(fmt.Errorf("Failed to read lstSearchCriteria from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.autoMatchmakeWithSearchCriteria_PostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.autoMatchmakeWithSearchCriteria_PostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	go protocol.autoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, lstSearchCriteria.([]*match_making_types.MatchmakeSessionSearchCriteria), anyGathering, strMessage)
}
