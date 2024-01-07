// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetDetailedList(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetDetailedList == nil {
		globals.Logger.Warning("Friends::GetDetailedList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	byRelationship := types.NewPrimitiveU8(0)
	err = byRelationship.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetDetailedList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReversed := types.NewPrimitiveBool(false)
	err = bReversed.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetDetailedList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetDetailedList(nil, packet, callID, byRelationship, bReversed)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
