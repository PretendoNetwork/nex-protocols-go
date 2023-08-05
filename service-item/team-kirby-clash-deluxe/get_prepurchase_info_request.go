// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetPrepurchaseInfoRequest sets the GetPrepurchaseInfoRequest handler function
func (protocol *Protocol) GetPrepurchaseInfoRequest(handler func(err error, client *nex.Client, callID uint32, getPrepurchaseInfoParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPrepurchaseInfoParam) uint32) {
	protocol.getPrepurchaseInfoRequestHandler = handler
}

func (protocol *Protocol) handleGetPrepurchaseInfoRequest(packet nex.PacketInterface) {
	if protocol.getPrepurchaseInfoRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPrepurchaseInfoRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getPrepurchaseInfoParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetPrepurchaseInfoParam())
	if err != nil {
		go protocol.getPrepurchaseInfoRequestHandler(fmt.Errorf("Failed to read getPrepurchaseInfoParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getPrepurchaseInfoRequestHandler(nil, client, callID, getPrepurchaseInfoParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetPrepurchaseInfoParam))
}
