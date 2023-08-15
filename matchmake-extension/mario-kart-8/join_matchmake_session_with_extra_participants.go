// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSessionWithExtraParticipants sets the JoinMatchmakeSessionWithExtraParticipants handler function
func (protocol *Protocol) JoinMatchmakeSessionWithExtraParticipants(handler func(err error, client *nex.Client, callID uint32, gid uint32, joinMessage string, ignoreBlacklist bool, participationCount uint16, extraParticipants uint32) uint32) {
	protocol.joinMatchmakeSessionWithExtraParticipantsHandler = handler
}

func (protocol *Protocol) handleJoinMatchmakeSessionWithExtraParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinMatchmakeSessionWithExtraParticipantsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	joinMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(fmt.Errorf("Failed to read joinMessage from parameters. %s", err.Error()), client, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	ignoreBlacklist, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(fmt.Errorf("Failed to read ignoreBlacklist from parameters. %s", err.Error()), client, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), client, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraParticipants, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(fmt.Errorf("Failed to read extraParticipants from parameters. %s", err.Error()), client, callID, 0, "", false, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.joinMatchmakeSessionWithExtraParticipantsHandler(nil, client, callID, gid, joinMessage, ignoreBlacklist, participationCount, extraParticipants)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
