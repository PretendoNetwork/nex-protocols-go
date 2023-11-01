// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// CreateCommunity sets the CreateCommunity handler function
func (protocol *Protocol) CreateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage string) uint32) {
	protocol.createCommunityHandler = handler
}

func (protocol *Protocol) handleCreateCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	community, err := parametersStream.ReadStructure(match_making_types.NewPersistentGathering())
	if err != nil {
		errorCode = protocol.createCommunityHandler(fmt.Errorf("Failed to read community from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createCommunityHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createCommunityHandler(nil, packet, callID, community.(*match_making_types.PersistentGathering), strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
