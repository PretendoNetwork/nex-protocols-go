// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPlayingSession sets the GetPlayingSession handler function
func (protocol *Protocol) GetPlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, lstPID []*nex.PID) uint32) {
	protocol.getPlayingSessionHandler = handler
}

func (protocol *Protocol) handleGetPlayingSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetPlayingSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPID, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.getPlayingSessionHandler(fmt.Errorf("Failed to read lstPID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPlayingSessionHandler(nil, packet, callID, lstPID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
