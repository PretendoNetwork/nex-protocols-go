// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetServiceItemRightRequest sets the GetServiceItemRightRequest handler function
func (protocol *Protocol) GetServiceItemRightRequest(handler func(err error, client *nex.Client, callID uint32, getServiceItemRightParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam, withoutRightBinary bool)) {
	protocol.getServiceItemRightRequestHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	if protocol.getServiceItemRightRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetServiceItemRightRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getServiceItemRightParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetServiceItemRightParam())
	if err != nil {
		go protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	withoutRightBinary, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read withoutRightBinary from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	go protocol.getServiceItemRightRequestHandler(nil, client, callID, getServiceItemRightParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam), withoutRightBinary)
}