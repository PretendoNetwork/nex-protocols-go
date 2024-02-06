// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindMatchmakeSessionByGatheringID(packet nex.PacketInterface) {
	if protocol.FindMatchmakeSessionByGatheringID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::FindMatchmakeSessionByGatheringID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstGID := types.NewList[*types.PrimitiveU32]()
	lstGID.Type = types.NewPrimitiveU32(0)

	err := lstGID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindMatchmakeSessionByGatheringID(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindMatchmakeSessionByGatheringID(nil, packet, callID, lstGID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
