// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmakeWithGatheringIDPostpone sets the AutoMatchmakeWithGatheringIDPostpone handler function
func (protocol *Protocol) AutoMatchmakeWithGatheringIDPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID []uint32, anyGathering *nex.DataHolder, strMessage string) uint32) {
	protocol.autoMatchmakeWithGatheringIDPostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakeWithGatheringIDPostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.autoMatchmakeWithGatheringIDPostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithGatheringIDPostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.autoMatchmakeWithGatheringIDPostponeHandler(nil, packet, callID, lstGID, anyGathering, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
