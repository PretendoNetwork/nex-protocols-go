// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Participate sets the Participate handler function
func (protocol *Protocol) Participate(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32) {
	protocol.participateHandler = handler
}

func (protocol *Protocol) handleParticipate(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.participateHandler == nil {
		globals.Logger.Warning("MatchMaking::Participate not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.participateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.participateHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.participateHandler(nil, packet, callID, idGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
