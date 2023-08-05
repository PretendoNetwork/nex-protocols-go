// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetServiceItemRightRequest sets the GetServiceItemRightRequest handler function
func (protocol *Protocol) GetServiceItemRightRequest(handler func(err error, client *nex.Client, callID uint32, getServiceItemRightParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam, withoutRightBinary bool) uint32) {
	protocol.getServiceItemRightRequestHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getServiceItemRightRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetServiceItemRightRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getServiceItemRightParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetServiceItemRightParam())
	if err != nil {
		errorCode = protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	withoutRightBinary, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getServiceItemRightRequestHandler(fmt.Errorf("Failed to read withoutRightBinary from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getServiceItemRightRequestHandler(nil, client, callID, getServiceItemRightParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam), withoutRightBinary)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
