// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleListServiceItemResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ListServiceItemResponse == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::ListServiceItemResponse not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ListServiceItemResponse(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ListServiceItemResponse(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
