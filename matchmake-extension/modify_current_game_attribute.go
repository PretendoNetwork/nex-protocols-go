// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleModifyCurrentGameAttribute(packet nex.PacketInterface) {
	var err error

	if protocol.ModifyCurrentGameAttribute == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::ModifyCurrentGameAttribute not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	attribIndex := types.NewPrimitiveU32(0)
	err = attribIndex.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read attribIndex from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	newValue := types.NewPrimitiveU32(0)
	err = newValue.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read newValue from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ModifyCurrentGameAttribute(nil, packet, callID, gid, attribIndex, newValue)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
