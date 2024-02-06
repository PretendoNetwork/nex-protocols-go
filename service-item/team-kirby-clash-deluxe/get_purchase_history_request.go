// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetPurchaseHistoryRequest(packet nex.PacketInterface) {
	if protocol.GetPurchaseHistoryRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemTeamKirbyClashDeluxe::GetPurchaseHistoryRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getPurchaseHistoryParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetPurchaseHistoryParam()

	err := getPurchaseHistoryParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPurchaseHistoryRequest(fmt.Errorf("Failed to read getPurchaseHistoryParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPurchaseHistoryRequest(nil, packet, callID, getPurchaseHistoryParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
