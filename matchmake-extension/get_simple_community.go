// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimpleCommunity sets the GetSimpleCommunity handler function
func (protocol *Protocol) GetSimpleCommunity(handler func(err error, client *nex.Client, callID uint32, gatheringIDList []uint32)) {
	protocol.getSimpleCommunityHandler = handler
}

func (protocol *Protocol) handleGetSimpleCommunity(packet nex.PacketInterface) {
	if protocol.getSimpleCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimpleCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gatheringIDList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getSimpleCommunityHandler(fmt.Errorf("Failed to read gatheringIDList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getSimpleCommunityHandler(nil, client, callID, gatheringIDList)
}
