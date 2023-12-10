// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionWithExtraParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.JoinMatchmakeSessionWithExtraParticipants == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	joinMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read joinMessage from parameters. %s", err.Error()), packet, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	ignoreBlacklist, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read ignoreBlacklist from parameters. %s", err.Error()), packet, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraParticipants, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionWithExtraParticipants(fmt.Errorf("Failed to read extraParticipants from parameters. %s", err.Error()), packet, callID, 0, "", false, 0, 0)
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
