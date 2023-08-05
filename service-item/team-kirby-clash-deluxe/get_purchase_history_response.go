// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPurchaseHistoryResponse sets the GetPurchaseHistoryResponse handler function
func (protocol *Protocol) GetPurchaseHistoryResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32) uint32) {
	protocol.getPurchaseHistoryResponseHandler = handler
}

func (protocol *Protocol) handleGetPurchaseHistoryResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPurchaseHistoryResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPurchaseHistoryResponse not implemented")
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
		errorCode = protocol.getPurchaseHistoryResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPurchaseHistoryResponseHandler(nil, client, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
