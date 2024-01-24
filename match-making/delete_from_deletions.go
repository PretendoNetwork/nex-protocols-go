// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteFromDeletions(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeleteFromDeletions == nil {
		globals.Logger.Warning("MatchMaking::DeleteFromDeletions not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstDeletions := types.NewList[*types.PrimitiveU32]()
	lstDeletions.Type = types.NewPrimitiveU32(0)
	err = lstDeletions.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteFromDeletions(fmt.Errorf("Failed to read lstDeletions from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteFromDeletions(nil, packet, callID, lstDeletions)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
