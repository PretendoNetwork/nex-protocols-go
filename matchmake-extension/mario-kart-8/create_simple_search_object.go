// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

func (protocol *Protocol) handleCreateSimpleSearchObject(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.CreateSimpleSearchObject == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::CreateSimpleSearchObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	object := matchmake_extension_mario_kart8_types.NewSimpleSearchObject()
	err = object.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateSimpleSearchObject(fmt.Errorf("Failed to read object from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CreateSimpleSearchObject(nil, packet, callID, object)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
