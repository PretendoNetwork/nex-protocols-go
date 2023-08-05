// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrepurchaseInfoResponse sets the GetPrepurchaseInfoResponse handler function
func (protocol *Protocol) GetPrepurchaseInfoResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32) uint32) {
	protocol.getPrepurchaseInfoResponseHandler = handler
}

func (protocol *Protocol) handleGetPrepurchaseInfoResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPrepurchaseInfoResponseHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetPrepurchaseInfoResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getPrepurchaseInfoResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPrepurchaseInfoResponseHandler(nil, client, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
