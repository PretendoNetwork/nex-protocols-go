// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CancelParticipation sets the CancelParticipation handler function
func (protocol *Protocol) CancelParticipation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string) uint32) {
	protocol.cancelParticipationHandler = handler
}

func (protocol *Protocol) handleCancelParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.cancelParticipationHandler == nil {
		globals.Logger.Warning("MatchMaking::CancelParticipation not implemented")
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
		errorCode = protocol.cancelParticipationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.cancelParticipationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.cancelParticipationHandler(nil, client, callID, idGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
