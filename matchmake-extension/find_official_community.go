package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindOfficialCommunity sets the FindOfficialCommunity handler function
func (protocol *MatchmakeExtensionProtocol) FindOfficialCommunity(handler func(err error, client *nex.Client, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange)) {
	protocol.FindOfficialCommunityHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleFindOfficialCommunity(packet nex.PacketInterface) {
	if protocol.FindOfficialCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindOfficialCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	isAvailableOnly := parametersStream.ReadBool()
	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.FindOfficialCommunityHandler(nil, client, callID, false, nil)
		return
	}

	go protocol.FindOfficialCommunityHandler(nil, client, callID, isAvailableOnly, resultRange.(*nex.ResultRange))
}
