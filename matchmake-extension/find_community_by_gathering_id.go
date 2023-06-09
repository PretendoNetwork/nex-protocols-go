package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByGatheringID sets the FindCommunityByGatheringID handler function
func (protocol *MatchmakeExtensionProtocol) FindCommunityByGatheringID(handler func(err error, client *nex.Client, callID uint32, lstGID []uint32)) {
	protocol.FindCommunityByGatheringIDHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleFindCommunityByGatheringID(packet nex.PacketInterface) {
	if protocol.FindCommunityByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByGatheringID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.FindCommunityByGatheringIDHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.FindCommunityByGatheringIDHandler(nil, client, callID, lstGID)
}
