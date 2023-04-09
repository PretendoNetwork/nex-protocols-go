package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimpleCommunity sets the GetSimpleCommunity handler function
func (protocol *MatchmakeExtensionProtocol) GetSimpleCommunity(handler func(err error, client *nex.Client, callID uint32, gatheringIDList []uint32)) {
	protocol.GetSimpleCommunityHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleGetSimpleCommunity(packet nex.PacketInterface) {
	if protocol.GetSimpleCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimpleCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gatheringIDList := parametersStream.ReadListUInt32LE()

	go protocol.GetSimpleCommunityHandler(nil, client, callID, gatheringIDList)
}
