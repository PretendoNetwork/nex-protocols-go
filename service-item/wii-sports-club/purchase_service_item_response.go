// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PurchaseServiceItemResponse sets the PurchaseServiceItemResponse handler function
func (protocol *Protocol) PurchaseServiceItemResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.purchaseServiceItemResponseHandler = handler
}

func (protocol *Protocol) handlePurchaseServiceItemResponse(packet nex.PacketInterface) {
	if protocol.purchaseServiceItemResponseHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::PurchaseServiceItemResponse not implemented")
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
		go protocol.purchaseServiceItemResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.purchaseServiceItemResponseHandler(nil, client, callID, requestID)
}