// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleCreateCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CreateCommunity == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	community, err := parametersStream.ReadStructure(match_making_types.NewPersistentGathering())
	if err != nil {
		_, errorCode = protocol.CreateCommunity(fmt.Errorf("Failed to read community from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.CreateCommunity(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CreateCommunity(nil, packet, callID, community.(*match_making_types.PersistentGathering), strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
