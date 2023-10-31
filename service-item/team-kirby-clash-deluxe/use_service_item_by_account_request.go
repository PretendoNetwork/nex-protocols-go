// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// UseServiceItemByAccountRequest sets the UseServiceItemByAccountRequest handler function
func (protocol *Protocol) UseServiceItemByAccountRequest(handler func(err error, packet nex.PacketInterface, callID uint32, useServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam) uint32) {
	protocol.useServiceItemByAccountRequestHandler = handler
}

func (protocol *Protocol) handleUseServiceItemByAccountRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.useServiceItemByAccountRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::UseServiceItemByAccountRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	useServiceItemByAccountParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemUseServiceItemByAccountParam())
	if err != nil {
		errorCode = protocol.useServiceItemByAccountRequestHandler(fmt.Errorf("Failed to read useServiceItemByAccountParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.useServiceItemByAccountRequestHandler(nil, packet, callID, useServiceItemByAccountParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
