// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProgressScore sets the UpdateProgressScore handler function
func (protocol *Protocol) UpdateProgressScore(handler func(err error, client *nex.Client, callID uint32, gid uint32, progressScore uint8) uint32) {
	protocol.updateProgressScoreHandler = handler
}

func (protocol *Protocol) handleUpdateProgressScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateProgressScoreHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateProgressScore not implemented")
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
		errorCode = protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	progressScore, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read progressScore from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateProgressScoreHandler(nil, client, callID, gid, progressScore)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
