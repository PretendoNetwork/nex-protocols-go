// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPurchaseHistoryResponse sets the GetPurchaseHistoryResponse handler function
func (protocol *Protocol) GetPurchaseHistoryResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.getPurchaseHistoryResponseHandler = handler
}

func (protocol *Protocol) handleGetPurchaseHistoryResponse(packet nex.PacketInterface) {
	if protocol.getPurchaseHistoryResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPurchaseHistoryResponse not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getPurchaseHistoryResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getPurchaseHistoryResponseHandler(nil, client, callID, requestID)
}
