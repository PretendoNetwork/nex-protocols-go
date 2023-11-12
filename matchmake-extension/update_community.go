// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// UpdateCommunity sets the UpdateCommunity handler function
func (protocol *Protocol) UpdateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) uint32) {
	protocol.updateCommunityHandler = handler
}

func (protocol *Protocol) handleUpdateCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	community, err := parametersStream.ReadStructure(match_making_types.NewPersistentGathering())
	if err != nil {
		errorCode = protocol.updateCommunityHandler(fmt.Errorf("Failed to read community from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateCommunityHandler(nil, packet, callID, community.(*match_making_types.PersistentGathering))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
