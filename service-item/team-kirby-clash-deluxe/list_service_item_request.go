// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// ListServiceItemRequest sets the ListServiceItemRequest handler function
func (protocol *Protocol) ListServiceItemRequest(handler func(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam) uint32) {
	protocol.listServiceItemRequestHandler = handler
}

func (protocol *Protocol) handleListServiceItemRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.listServiceItemRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::ListServiceItemRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listServiceItemParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemListServiceItemParam())
	if err != nil {
		errorCode = protocol.listServiceItemRequestHandler(fmt.Errorf("Failed to read listServiceItemParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.listServiceItemRequestHandler(nil, packet, callID, listServiceItemParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
