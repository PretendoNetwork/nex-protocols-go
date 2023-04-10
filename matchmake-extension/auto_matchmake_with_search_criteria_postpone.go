package matchmake_extension

import (
	"encoding/hex"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// AutoMatchmakeWithSearchCriteria_Postpone sets the AutoMatchmakeWithSearchCriteria_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmakeWithSearchCriteria_Postpone(handler func(err error, client *nex.Client, callID uint32, lstSearchCriteria []*match_making.MatchmakeSessionSearchCriteria, anyGathering *nex.DataHolder, strMessage string)) {
	protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleAutoMatchmakeWithSearchCriteria_Postpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler == nil {
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

	lstSearchCriteria, err := parametersStream.ReadListStructure(match_making.NewMatchmakeSessionSearchCriteria())
	if err != nil {
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, nil, nil, "")
	}

	anyGathering := parametersStream.ReadDataHolder()
	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, nil, nil, "")
	}

	go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, lstSearchCriteria.([]*match_making.MatchmakeSessionSearchCriteria), anyGathering, strMessage)
}
