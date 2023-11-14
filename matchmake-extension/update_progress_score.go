// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateProgressScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateProgressScore == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateProgressScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.UpdateProgressScore(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	progressScore, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.UpdateProgressScore(fmt.Errorf("Failed to read progressScore from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdateProgressScore(nil, packet, callID, gid, progressScore)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
