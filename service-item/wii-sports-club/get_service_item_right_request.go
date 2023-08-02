// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetServiceItemRightRequest sets the GetServiceItemRightRequest handler function
func (protocol *Protocol) GetServiceItemRightRequest(handler func(err error, client *nex.Client, callID uint32, getServiceItemRightParam *service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam)) {
	protocol.getServiceItemRightRequestHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	if protocol.getServiceItemRightRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetServiceItemRightRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getServiceItemRightParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemGetServiceItemRightParam())
	if err != nil {
		go protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getServiceItemRightRequestHandler(nil, client, callID, getServiceItemRightParam.(*service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam))
}