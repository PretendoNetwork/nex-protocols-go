// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// GetServiceItemRightRequest sets the GetServiceItemRightRequest handler function
func (protocol *Protocol) GetServiceItemRightRequest(handler func(err error, packet nex.PacketInterface, callID uint32, getServiceItemRightParam *service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam) uint32) {
	protocol.getServiceItemRightRequestHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getServiceItemRightRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetServiceItemRightRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getServiceItemRightParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemGetServiceItemRightParam())
	if err != nil {
		errorCode = protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getServiceItemRightRequestHandler(nil, packet, callID, getServiceItemRightParam.(*service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
