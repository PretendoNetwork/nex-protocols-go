// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

func (protocol *Protocol) handleCreateSimpleSearchObject(packet nex.PacketInterface) {
	if protocol.CreateSimpleSearchObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMarioKart8::CreateSimpleSearchObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	object := matchmake_extension_mario_kart8_types.NewSimpleSearchObject()

	err := object.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateSimpleSearchObject(fmt.Errorf("Failed to read object from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateSimpleSearchObject(nil, packet, callID, object)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
