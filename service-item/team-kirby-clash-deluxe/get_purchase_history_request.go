// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetPurchaseHistoryRequest sets the GetPurchaseHistoryRequest handler function
func (protocol *Protocol) GetPurchaseHistoryRequest(handler func(err error, client *nex.Client, callID uint32, getPurchaseHistoryParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPurchaseHistoryParam) uint32) {
	protocol.getPurchaseHistoryRequestHandler = handler
}

func (protocol *Protocol) handleGetPurchaseHistoryRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPurchaseHistoryRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPurchaseHistoryRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getPurchaseHistoryParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetPurchaseHistoryParam())
	if err != nil {
		errorCode = protocol.getPurchaseHistoryRequestHandler(fmt.Errorf("Failed to read getPurchaseHistoryParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPurchaseHistoryRequestHandler(nil, client, callID, getPurchaseHistoryParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetPurchaseHistoryParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
