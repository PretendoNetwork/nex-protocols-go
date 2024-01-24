// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetList(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetList == nil {
		globals.Logger.Warning("Friends::GetList not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	byRelationship := types.NewPrimitiveU8(0)
	err = byRelationship.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReversed := types.NewPrimitiveBool(false)
	err = bReversed.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetList(nil, packet, callID, byRelationship, bReversed)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
