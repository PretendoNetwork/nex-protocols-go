// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

// SearchSimpleSearchObject sets the SearchSimpleSearchObject handler function
func (protocol *Protocol) SearchSimpleSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_extension_mario_kart8_types.SimpleSearchParam) uint32) {
	protocol.searchSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleSearchSimpleSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchSimpleSearchObjectHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::SearchSimpleSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_extension_mario_kart8_types.NewSimpleSearchParam())
	if err != nil {
		errorCode = protocol.searchSimpleSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.searchSimpleSearchObjectHandler(nil, packet, callID, param.(*matchmake_extension_mario_kart8_types.SimpleSearchParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
