// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStats sets the GetStats handler function
func (protocol *Protocol) GetStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstParticipants []*nex.PID, lstColumns []byte) uint32) {
	protocol.getStatsHandler = handler
}

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStatsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstParticipants, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read lstParticipants from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstColumns, err := parametersStream.ReadBuffer() // * This is documented as List<byte>, but that's justs a buffer so...
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read lstColumns from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStatsHandler(nil, packet, callID, idGathering, lstParticipants, lstColumns)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
