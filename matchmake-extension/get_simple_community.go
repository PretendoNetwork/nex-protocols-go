// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimpleCommunity sets the GetSimpleCommunity handler function
func (protocol *Protocol) GetSimpleCommunity(handler func(err error, client *nex.Client, callID uint32, gatheringIDList []uint32) uint32) {
	protocol.getSimpleCommunityHandler = handler
}

func (protocol *Protocol) handleGetSimpleCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSimpleCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimpleCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gatheringIDList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getSimpleCommunityHandler(fmt.Errorf("Failed to read gatheringIDList from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSimpleCommunityHandler(nil, client, callID, gatheringIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
