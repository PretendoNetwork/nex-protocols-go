// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAutoMatchmakeWithGatheringIDPostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AutoMatchmakeWithGatheringIDPostpone == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithGatheringIDPostpone not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AutoMatchmakeWithGatheringIDPostpone(nil, packet, callID, lstGID, anyGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
