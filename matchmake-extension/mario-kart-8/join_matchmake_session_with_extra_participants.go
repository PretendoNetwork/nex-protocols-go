// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionWithExtraParticipants(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.JoinMatchmakeSessionWithExtraParticipants == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	joinMessage := types.NewString("")
	err = joinMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read joinMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	ignoreBlacklist := types.NewPrimitiveBool(false)
	err = ignoreBlacklist.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read ignoreBlacklist from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount := types.NewPrimitiveU16(0)
	err = participationCount.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraParticipants := types.NewPrimitiveU32(0)
	err = extraParticipants.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read extraParticipants from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.JoinMatchmakeSessionWithExtraParticipants(nil, packet, callID, gid, joinMessage, ignoreBlacklist, participationCount, extraParticipants)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
