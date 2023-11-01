// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// PurchaseServiceItemRequest sets the PurchaseServiceItemRequest handler function
func (protocol *Protocol) PurchaseServiceItemRequest(handler func(err error, packet nex.PacketInterface, callID uint32, purchaseServiceItemParam *service_item_wii_sports_club_types.ServiceItemPurchaseServiceItemParam) uint32) {
	protocol.purchaseServiceItemRequestHandler = handler
}

func (protocol *Protocol) handlePurchaseServiceItemRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.purchaseServiceItemRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::PurchaseServiceItemRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	purchaseServiceItemParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemPurchaseServiceItemParam())
	if err != nil {
		errorCode = protocol.purchaseServiceItemRequestHandler(fmt.Errorf("Failed to read purchaseServiceItemParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.purchaseServiceItemRequestHandler(nil, packet, callID, purchaseServiceItemParam.(*service_item_wii_sports_club_types.ServiceItemPurchaseServiceItemParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
