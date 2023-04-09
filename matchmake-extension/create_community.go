package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// CreateCommunity sets the CreateCommunity handler function
func (protocol *MatchmakeExtensionProtocol) CreateCommunity(handler func(err error, client *nex.Client, callID uint32, community *match_making.PersistentGathering, strMessage string)) {
	protocol.CreateCommunityHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleCreateCommunity(packet nex.PacketInterface) {
	if protocol.CreateCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	community, err := parametersStream.ReadStructure(match_making.NewPersistentGathering())
	if err != nil {
		go protocol.CreateCommunityHandler(nil, client, callID, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.CreateCommunityHandler(nil, client, callID, nil, "")
		return
	}

	go protocol.CreateCommunityHandler(nil, client, callID, community.(*match_making.PersistentGathering), strMessage)
}
