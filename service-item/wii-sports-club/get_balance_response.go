// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBalanceResponse sets the GetBalanceResponse handler function
func (protocol *Protocol) GetBalanceResponse(handler func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32) {
	protocol.getBalanceResponseHandler = handler
}

func (protocol *Protocol) handleGetBalanceResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getBalanceResponseHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetBalanceResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getBalanceResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getBalanceResponseHandler(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
