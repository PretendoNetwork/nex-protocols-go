// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmakePostpone sets the AutoMatchmakePostpone handler function
func (protocol *Protocol) AutoMatchmakePostpone(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder, strMessage string) uint32) {
	protocol.autoMatchmakePostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakePostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.autoMatchmakePostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakePostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.autoMatchmakePostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.autoMatchmakePostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.autoMatchmakePostponeHandler(nil, packet, callID, anyGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
