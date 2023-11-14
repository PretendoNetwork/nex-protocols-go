// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

func (protocol *Protocol) handleCreateSimpleSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CreateSimpleSearchObject == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::CreateSimpleSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	object, err := parametersStream.ReadStructure(matchmake_extension_mario_kart8_types.NewSimpleSearchObject())
	if err != nil {
		errorCode = protocol.CreateSimpleSearchObject(fmt.Errorf("Failed to read object from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.CreateSimpleSearchObject(nil, packet, callID, object.(*matchmake_extension_mario_kart8_types.SimpleSearchObject))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
