// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByNameLike(packet nex.PacketInterface) {
	var err error

	if protocol.FindByNameLike == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::FindByNameLike not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroups := types.NewPrimitiveU32(0)
	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strLike := types.NewString("")
	err = strLike.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	resultRange := types.NewResultRange()
	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByNameLike(nil, packet, callID, uiGroups, strLike, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
