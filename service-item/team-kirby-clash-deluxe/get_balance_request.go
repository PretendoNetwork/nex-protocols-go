// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBalanceRequest sets the GetBalanceRequest handler function
func (protocol *Protocol) GetBalanceRequest(handler func(err error, client *nex.Client, callID uint32, getBalanceParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetBalanceParam)) {
	protocol.getBalanceRequestHandler = handler
}

func (protocol *Protocol) handleGetBalanceRequest(packet nex.PacketInterface) {
	if protocol.getBalanceRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetBalanceRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getBalanceParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetBalanceParam())
	if err != nil {
		go protocol.getBalanceRequestHandler(fmt.Errorf("Failed to read getBalanceParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getBalanceRequestHandler(nil, client, callID, getBalanceParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetBalanceParam))
}