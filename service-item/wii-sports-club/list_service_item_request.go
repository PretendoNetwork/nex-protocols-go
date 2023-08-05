// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// ListServiceItemRequest sets the ListServiceItemRequest handler function
func (protocol *Protocol) ListServiceItemRequest(handler func(err error, client *nex.Client, callID uint32, listServiceItemParam *service_item_wii_sports_club_types.ServiceItemListServiceItemParam) uint32) {
	protocol.listServiceItemRequestHandler = handler
}

func (protocol *Protocol) handleListServiceItemRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.listServiceItemRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::ListServiceItemRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listServiceItemParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemListServiceItemParam())
	if err != nil {
		errorCode = protocol.listServiceItemRequestHandler(fmt.Errorf("Failed to read listServiceItemParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.listServiceItemRequestHandler(nil, client, callID, listServiceItemParam.(*service_item_wii_sports_club_types.ServiceItemListServiceItemParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
