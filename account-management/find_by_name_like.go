// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByNameLike(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindByNameLike == nil {
		globals.Logger.Warning("AccountManagement::FindByNameLike not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strLike, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := nex.StreamReadStructure(parametersStream, nex.NewResultRange())
	if err != nil {
		_, errorCode = protocol.FindByNameLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, "", nil)
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
