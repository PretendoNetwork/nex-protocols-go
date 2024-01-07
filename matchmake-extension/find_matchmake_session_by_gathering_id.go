// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindMatchmakeSessionByGatheringID(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.FindMatchmakeSessionByGatheringID == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByGatheringID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstGID := types.NewList[*types.PrimitiveU32]()
	lstGID.Type = types.NewPrimitiveU32(0)
	err = lstGID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindMatchmakeSessionByGatheringID(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindMatchmakeSessionByGatheringID(nil, packet, callID, lstGID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
