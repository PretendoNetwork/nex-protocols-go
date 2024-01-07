// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handlePurchaseServiceItemRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.PurchaseServiceItemRequest == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::PurchaseServiceItemRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	purchaseServiceItemParam := service_item_wii_sports_club_types.NewServiceItemPurchaseServiceItemParam()
	err = purchaseServiceItemParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PurchaseServiceItemRequest(fmt.Errorf("Failed to read purchaseServiceItemParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PurchaseServiceItemRequest(nil, packet, callID, purchaseServiceItemParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
