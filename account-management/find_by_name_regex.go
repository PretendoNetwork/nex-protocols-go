// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByNameRegex(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.FindByNameRegex == nil {
		globals.Logger.Warning("AccountManagement::FindByNameRegex not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroups := types.NewPrimitiveU32(0)
	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strRegex := types.NewString("")
	err = strRegex.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read strRegex from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange := types.NewResultRange()
	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByNameRegex(nil, packet, callID, uiGroups, strRegex, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
