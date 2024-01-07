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
	var errorCode uint32

	if protocol.FindByNameLike == nil {
		globals.Logger.Warning("AccountManagement::FindByNameLike not implemented")
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
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strLike := types.NewString("")
	err = strLike.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange := types.NewResultRange()
	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByNameLike(nil, packet, callID, uiGroups, strLike, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
