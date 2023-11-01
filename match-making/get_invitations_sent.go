// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetInvitationsSent sets the GetInvitationsSent handler function
func (protocol *Protocol) GetInvitationsSent(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32) {
	protocol.getInvitationsSentHandler = handler
}

func (protocol *Protocol) handleGetInvitationsSent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getInvitationsSentHandler == nil {
		globals.Logger.Warning("MatchMaking::GetInvitationsSent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getInvitationsSentHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getInvitationsSentHandler(nil, packet, callID, idGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
