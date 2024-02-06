// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSimpleCommunity(packet nex.PacketInterface) {
	if protocol.GetSimpleCommunity == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::GetSimpleCommunity not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gatheringIDList := types.NewList[*types.PrimitiveU32]()
	gatheringIDList.Type = types.NewPrimitiveU32(0)

	err := gatheringIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSimpleCommunity(fmt.Errorf("Failed to read gatheringIDList from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSimpleCommunity(nil, packet, callID, gatheringIDList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
