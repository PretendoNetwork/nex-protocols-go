// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// ListServiceItemRequest sets the ListServiceItemRequest handler function
func (protocol *Protocol) ListServiceItemRequest(handler func(err error, client *nex.Client, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam)) {
	protocol.listServiceItemRequestHandler = handler
}

func (protocol *Protocol) handleListServiceItemRequest(packet nex.PacketInterface) {
	if protocol.listServiceItemRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::ListServiceItemRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listServiceItemParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemListServiceItemParam())
	if err != nil {
		go protocol.listServiceItemRequestHandler(fmt.Errorf("Failed to read listServiceItemParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.listServiceItemRequestHandler(nil, client, callID, listServiceItemParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam))
}
