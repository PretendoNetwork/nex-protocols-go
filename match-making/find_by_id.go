// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByID(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.FindByID == nil {
		globals.Logger.Warning("MatchMaking::FindByID not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstID := types.NewList[*types.PrimitiveU32]()
	lstID.Type = types.NewPrimitiveU32(0)
	err = lstID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByID(fmt.Errorf("Failed to read lstID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByID(nil, packet, callID, lstID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
