// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByParticipant sets the FindCommunityByParticipant handler function
func (protocol *Protocol) FindCommunityByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, pid uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findCommunityByParticipantHandler = handler
}

func (protocol *Protocol) handleFindCommunityByParticipant(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findCommunityByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByParticipant not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findCommunityByParticipantHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findCommunityByParticipantHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findCommunityByParticipantHandler(nil, packet, callID, pid, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
