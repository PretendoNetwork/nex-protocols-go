// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetDetailedList(packet nex.PacketInterface) {
	if protocol.GetDetailedList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::GetDetailedList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	byRelationship := types.NewPrimitiveU8(0)
	bReversed := types.NewPrimitiveBool(false)

	var err error

	err = byRelationship.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetDetailedList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = bReversed.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetDetailedList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetDetailedList(nil, packet, callID, byRelationship, bReversed)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
