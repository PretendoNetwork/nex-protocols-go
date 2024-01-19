// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetServiceItemRightResponse(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetServiceItemRightResponse == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetServiceItemRightResponse not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	requestID := types.NewPrimitiveU32(0)
	err = requestID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetServiceItemRightResponse(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetServiceItemRightResponse(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
