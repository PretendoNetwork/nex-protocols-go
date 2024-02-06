// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMatchmakeSessionAttribute(packet nex.PacketInterface) {
	if protocol.UpdateMatchmakeSessionAttribute == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateMatchmakeSessionAttribute not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	attribs := types.NewList[*types.PrimitiveU32]()
	attribs.Type = types.NewPrimitiveU32(0)

	var err error

	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMatchmakeSessionAttribute(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = attribs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMatchmakeSessionAttribute(fmt.Errorf("Failed to read attribs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMatchmakeSessionAttribute(nil, packet, callID, gid, attribs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
