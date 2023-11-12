// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDetailedParticipants sets the GetDetailedParticipants handler function
func (protocol *Protocol) GetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, bOnlyActive bool) uint32) {
	protocol.getDetailedParticipantsHandler = handler
}

func (protocol *Protocol) handleGetDetailedParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getDetailedParticipantsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetDetailedParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getDetailedParticipantsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bOnlyActive, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getDetailedParticipantsHandler(fmt.Errorf("Failed to read bOnlyActive from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getDetailedParticipantsHandler(nil, packet, callID, idGathering, bOnlyActive)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
