// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByNameRegex(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindByNameRegex == nil {
		globals.Logger.Warning("AccountManagement::FindByNameRegex not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strRegex, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read strRegex from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.FindByNameRegex(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.FindByNameRegex(nil, packet, callID, uiGroups, strRegex, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
