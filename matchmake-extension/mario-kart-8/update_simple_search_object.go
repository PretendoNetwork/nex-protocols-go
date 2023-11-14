// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

func (protocol *Protocol) handleUpdateSimpleSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateSimpleSearchObject == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::UpdateSimpleSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.UpdateSimpleSearchObject(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	newObject, err := parametersStream.ReadStructure(matchmake_extension_mario_kart8_types.NewSimpleSearchObject())
	if err != nil {
		errorCode = protocol.UpdateSimpleSearchObject(fmt.Errorf("Failed to read newObject from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdateSimpleSearchObject(nil, packet, callID, objectID, newObject.(*matchmake_extension_mario_kart8_types.SimpleSearchObject))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
