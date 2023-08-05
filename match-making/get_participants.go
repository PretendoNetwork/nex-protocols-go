// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipants sets the GetParticipants handler function
func (protocol *Protocol) GetParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32) uint32) {
	protocol.getParticipantsHandler = handler
}

func (protocol *Protocol) handleGetParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getParticipantsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getParticipantsHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getParticipantsHandler(nil, client, callID, idGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
