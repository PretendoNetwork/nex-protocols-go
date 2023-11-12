// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteScore sets the DeleteScore handler function
func (protocol *Protocol) DeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, uniqueID uint64) uint32) {
	protocol.deleteScoreHandler = handler
}

func (protocol *Protocol) handleDeleteScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteScoreHandler == nil {
		globals.Logger.Warning("Ranking::DeleteScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.deleteScoreHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.deleteScoreHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteScoreHandler(nil, packet, callID, category, uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
