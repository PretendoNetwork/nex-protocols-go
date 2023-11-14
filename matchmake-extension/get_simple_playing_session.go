// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (protocol *Protocol) GetSimplePlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, listPID []*nex.PID, includeLoginUser bool) uint32) {
	protocol.getSimplePlayingSessionHandler = handler
}

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSimplePlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimplePlayingSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listPID, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.getSimplePlayingSessionHandler(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	includeLoginUser, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getSimplePlayingSessionHandler(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSimplePlayingSessionHandler(nil, packet, callID, listPID, includeLoginUser)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
