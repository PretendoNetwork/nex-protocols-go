// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAutoMatchmakePostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AutoMatchmakePostpone == nil {
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
		_, errorCode = protocol.AutoMatchmakePostpone(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.AutoMatchmakePostpone(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AutoMatchmakePostpone(nil, packet, callID, anyGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
