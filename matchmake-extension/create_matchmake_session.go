// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateMatchmakeSession sets the CreateMatchmakeSession handler function
func (protocol *Protocol) CreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder, message string, participationCount uint16) uint32) {
	protocol.createMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	var errorCode uint32

	if protocol.createMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	var participationCount uint16 = 0

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		participationCount, err = parametersStream.ReadUInt16LE()
		if err != nil {
			errorCode = protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read message from participationCount. %s", err.Error()), packet, callID, nil, "", 0)
			if errorCode != 0 {
				globals.RespondError(packet, ProtocolID, errorCode)
			}

			return
		}
	}

	errorCode = protocol.createMatchmakeSessionHandler(nil, packet, callID, anyGathering, message, participationCount)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
