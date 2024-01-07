// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSimpleCommunity(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetSimpleCommunity == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimpleCommunity not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gatheringIDList := types.NewList[*types.PrimitiveU32]()
	gatheringIDList.Type = types.NewPrimitiveU32(0)
	err = gatheringIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetSimpleCommunity(fmt.Errorf("Failed to read gatheringIDList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetSimpleCommunity(nil, packet, callID, gatheringIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
