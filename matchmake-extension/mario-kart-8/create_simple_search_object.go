// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

// CreateSimpleSearchObject sets the CreateSimpleSearchObject handler function
func (protocol *Protocol) CreateSimpleSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, object *matchmake_extension_mario_kart8_types.SimpleSearchObject) uint32) {
	protocol.createSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleCreateSimpleSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createSimpleSearchObjectHandler == nil {
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
		errorCode = protocol.createSimpleSearchObjectHandler(fmt.Errorf("Failed to read object from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createSimpleSearchObjectHandler(nil, packet, callID, object.(*matchmake_extension_mario_kart8_types.SimpleSearchObject))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
