// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionWithExtraParticipants(packet nex.PacketInterface) {
	if protocol.JoinMatchmakeSessionWithExtraParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	joinMessage := types.NewString("")
	ignoreBlacklist := types.NewPrimitiveBool(false)
	participationCount := types.NewPrimitiveU16(0)
	extraParticipants := types.NewPrimitiveU32(0)

	var err error

	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = joinMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read joinMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = ignoreBlacklist.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read ignoreBlacklist from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = participationCount.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = extraParticipants.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read extraParticipants from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.JoinMatchmakeSessionWithExtraParticipants(nil, packet, callID, gid, joinMessage, ignoreBlacklist, participationCount, extraParticipants)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
