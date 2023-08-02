// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PurchaseServiceItemRequest sets the PurchaseServiceItemRequest handler function
func (protocol *Protocol) PurchaseServiceItemRequest(handler func(err error, client *nex.Client, callID uint32, purchaseServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPurchaseServiceItemParam)) {
	protocol.purchaseServiceItemRequestHandler = handler
}

func (protocol *Protocol) handlePurchaseServiceItemRequest(packet nex.PacketInterface) {
	if protocol.purchaseServiceItemRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::PurchaseServiceItemRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	purchaseServiceItemParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemPurchaseServiceItemParam())
	if err != nil {
		go protocol.purchaseServiceItemRequestHandler(fmt.Errorf("Failed to read purchaseServiceItemParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.purchaseServiceItemRequestHandler(nil, client, callID, purchaseServiceItemParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemPurchaseServiceItemParam))
}